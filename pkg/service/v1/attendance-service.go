package v1

import (
	"context"
	"encoding/json"
	"github.com/djamboe/mtools-attendance-service/models"
	"github.com/streadway/amqp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"

	v1 "github.com/djamboe/mtools-attendance-service/pkg/api/v1"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

func init() {
	initAmqp()
}

var conn *amqp.Connection
var ch *amqp.Channel

func initAmqp() {
	var err error

	conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err = conn.Channel()
	failOnError(err, "Failed to open a channel")

}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// toDoServiceServer is implementation of v1.ToDoServiceServer proto interface
type attendanceServiceServer struct {
}

// NewToDoServiceServer creates ToDo service
func NewAttendanceServiceServer() v1.AttendanceServiceServer {
	return &attendanceServiceServer{}
}

// checkAPI checks if the API version requested by client is supported by server
func (s *attendanceServiceServer) checkAPI(api string) error {
	// API version is "" means use current version of the service
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	} else {
		return status.Errorf(codes.Unimplemented,
			"please input your api version")
	}
	return nil
}

// Create new todo task
func (s *attendanceServiceServer) CreateAttendance(ctx context.Context, req *v1.AttendanceRequest) (*v1.AttendanceResponse, error) {
	// check if the API version requested by client is supported by server
	var locationParam models.AttendanceParamModel
	message := "Successfully Create Location !"
	errorStatus := false
	if err := s.checkAPI(req.Api); err != nil {
		message = "Unsupported api version !"
		errorStatus = true
	}
	locationParam.UserId = req.UserId
	locationParam.Lat = req.Lat
	locationParam.Long = req.Long
	locationParam.Description = req.Description
	locationParam.MemberName = req.MemberName
	locationParam.ImageUrl = req.ImageUrl

	//locationController := ServiceContainer().InjectAttendanceController()
	//err := locationController.InsertNewLocation(locationParam)

	payload, err := json.Marshal(locationParam)
	failOnError(err, "Failed to marshal JSON")
	//try to publish message into broker
	err = ch.Publish(
		"post-attendance-exchange", // exchange
		"go-test-key-attendance",   // routing key
		false,                      // mandatory
		false,                      // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Transient,
			ContentType:  "application/json",
			Body:         payload,
			Timestamp:    time.Now(),
		})

	if err != nil {
		message = "Failed insert location, an error occured"
		errorStatus = true
	}

	return &v1.AttendanceResponse{
		Api:     apiVersion,
		Message: message,
		Error:   errorStatus,
	}, nil
}
