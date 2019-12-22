package v1

import (
	"context"
	"github.com/djamboe/mtools-attendance-service/controllers"
	"github.com/djamboe/mtools-attendance-service/infrastructures"
	"github.com/djamboe/mtools-attendance-service/repositories"
	"github.com/djamboe/mtools-attendance-service/services"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"sync"
)

type IserviceContainer interface {
	InjectAttendanceController() controllers.AttendanceController
}

type kernel struct{}

func (k *kernel) InjectAttendanceController() controllers.AttendanceController {
	//mysqlConn, _ := sql.Open("mysql", "root:@tcp(localhost:3306)/marketing-tools?charset=utf8")
	//mysqlHandler := &infrastructures.DBHandler{}
	//mysqlHandler.Conn = mysqlConn
	//loginRepository := &repositories.LoginRepository{mysqlHandler}
	//loginService := &services.LoginService{&repositories.LoginRepositoryWithCircuitBreaker{loginRepository}}
	//loginController := controllers.LoginController{loginService}

	//test mongodb
	c := GetClient()
	mongoCon := c.Ping(context.Background(), readpref.Primary())
	if mongoCon != nil {
		log.Fatal("Couldn't connect to the database", mongoCon)
	}

	mongoDBConn := c
	mongoDBHandler := &infrastructures.MongoDBHandler{}
	mongoDBHandler.Conn = mongoDBConn
	attendanceRepository := &repositories.AttendanceRepository{mongoDBHandler}
	attendanceService := &services.AttendanceService{&repositories.AttendanceRepositoryWithCircuitBreaker{attendanceRepository}}
	attendanceController := controllers.AttendanceController{attendanceService}
	//test mongodb

	return attendanceController
}

var (
	k             *kernel
	containerOnce sync.Once
)

func ServiceContainer() IserviceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}

func GetClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return client
}
