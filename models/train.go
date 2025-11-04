package models

import "time"

type Train struct {
	ID             string
	CurrentSegment string
	LastUpdated    time.Time
}
