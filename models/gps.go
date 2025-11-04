package models

import (
	"time"
)

type GPSData struct {
	TrainID   string
	Latitude  float64
	Longitude float64
	Timestamp time.Time
}
