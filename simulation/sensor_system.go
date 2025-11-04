// Now the Sensor System listens for GPS updates,
// checks which segment the coordinates belong to, and
// emits SensorEvents to the signalCh.
package simulation

import (
	"fmt"
	"math"
	"signal-system4/models"
)

// Helper: calculate distance between coordinates
func distance(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371 // km
	dLat := (lat2 - lat1) * math.Pi / 180
	dLon := (lon2 - lon1) * math.Pi / 180
	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1*math.Pi/180)*math.Cos(lat2*math.Pi/180)*
			math.Sin(dLon/2)*math.Sin(dLon/2)

	return R * 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
}

type SensorSystem struct{ Sensors []*models.Sensor }

func NewSensorSystem(segments []*models.Segment) *SensorSystem {
	sensors := make([]*models.Sensor, len(segments))
	for i, seg := range segments {
		sensors[i] = &models.Sensor{
			ID:      fmt.Sprintf("Sensor-%d", i),
			Segment: seg,
		}
	}

	return &SensorSystem{Sensors: sensors}
}

// ListenForGPS listens to GPS updates and triggers sensor events
func (ss *SensorSystem) ListenForGPS(gpsCh <-chan *models.GPSData, signalCh chan<- *models.SensorEvent) {
	const threshold = 0.15 // km threshold for segment detection
	for gps := range gpsCh {
		for _, sensor := range ss.Sensors {
			dist := distance(gps.Latitude, gps.Longitude,
				sensor.Segment.Latitude, sensor.Segment.Longitude)
			if dist < threshold {
				fmt.Printf("[SENSOR] Train %s ENTERED %s\n", gps.TrainID, sensor.Segment.Name)
				signalCh <- &models.SensorEvent{
					SensorID: sensor.ID,
					TrainID:  gps.TrainID,
					Type:     models.TrainEntered,
					Segment:  sensor.Segment,
				}
			}
		}
	}
}
