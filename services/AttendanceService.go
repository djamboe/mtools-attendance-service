package services

import (
	"github.com/djamboe/mtools-attendance-service/interfaces"
	"github.com/djamboe/mtools-attendance-service/models"
)

type AttendanceService struct {
	interfaces.IAttendanceRepository
}

func (service *AttendanceService) CreateAttendance(param models.AttendanceParamModel) error {
	err := service.InsertNewLocation(param)
	if err != nil {
		panic(err)
	}
	return nil
}
