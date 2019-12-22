package repositories

import (
	"github.com/afex/hystrix-go/hystrix"
	"github.com/djamboe/mtools-attendance-service/interfaces"
	"github.com/djamboe/mtools-attendance-service/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type AttendanceRepositoryWithCircuitBreaker struct {
	AttendanceRepository interfaces.IAttendanceRepository
}

type AttendanceRepository struct {
	//interfaces.IDbHandler
	interfaces.IMongoDBHandler
}

func (repository *AttendanceRepositoryWithCircuitBreaker) InsertNewLocation(param models.AttendanceParamModel) error {
	output := make(chan error, 1)
	hystrix.ConfigureCommand("create_post", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("create_post", func() error {
		locationData := repository.AttendanceRepository.InsertNewLocation(param)
		output <- locationData
		return nil
	}, nil)

	select {
	case out := <-output:
		return out
	case err := <-errors:
		println(err)
		return err
	}
}

func (repository *AttendanceRepository) InsertNewLocation(param models.AttendanceParamModel) error {
	var p models.Point
	p.UserId = param.UserId
	p.Location = models.NewPoint(param.Lat, param.Long)
	p.MemberName = param.MemberName
	p.Description = param.Description
	p.PresentTime = time.Now()
	p.ImageUrl = param.ImageUrl
	err := repository.InsertOne(p, "attendance", "maroon_martools")
	if err != nil {
		panic(err)
	}
	return nil
}
