package controllers

import (
	"github.com/djamboe/mtools-attendance-service/interfaces"
	"github.com/djamboe/mtools-attendance-service/models"
)

type AttendanceController struct {
	interfaces.IAttendanceRepository
}

func (controller *AttendanceController) RecordLocation(param models.AttendanceParamModel) error {
	err := controller.InsertNewLocation(param)
	if err != nil {
		panic(err)
	}
	return nil
}
