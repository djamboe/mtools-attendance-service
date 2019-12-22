package interfaces

import (
	"github.com/djamboe/mtools-attendance-service/models"
)

type IAttendanceRepository interface {
	InsertNewLocation(model models.AttendanceParamModel) error
}
