package models

// Point is a simple type with a location for geospatial
// queries.
type AttendanceResponseModel struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}
