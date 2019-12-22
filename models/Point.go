package models

import "time"

// Point is a simple type with a location for geospatial
// queries.
type Point struct {
	UserId      string    `json:"userid"`
	Description string    `json:"description"`
	MemberName  string    `json:"member_name"`
	Location    Location  `json:"location"`
	PresentTime time.Time `json:"presenttime"`
	ImageUrl    string    `json:"imageurl"`
}
