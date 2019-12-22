package models

// Point is a simple type with a location for geospatial
// queries.
type AttendanceParamModel struct {
	UserId      string  `json:"user_id"`
	Lat         float64 `json:"lat"`
	Long        float64 `json:"long"`
	Description string  `json:"description"`
	MemberName  string  `json:"membername"`
	ImageUrl    string  `json:"imageurl"`
}
