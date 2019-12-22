package interfaces

import "github.com/djamboe/mtools-attendance-service/models"

type ILocationService interface {
	CreateAttendance(param models.AttendanceParamModel) error
}
