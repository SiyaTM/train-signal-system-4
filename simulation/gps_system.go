// This simulates a train emitting position data over time through a channel.
package simulation

import (
	"fmt"
	"signal-system4/models"
	"time"
)

func StartGPSUpdates(train *models.Train, gpsCh chan<- *models.GPSData) {
	coords := [][]float64{
		{-25.755, 28.231},
		// segment A
		{-25.756, 28.232},
		// segment B
		{-25.757, 28.233},
		// segment C
	}

	for _, c := range coords {
		data := &models.GPSData{
			TrainID:   train.ID,
			Latitude:  c[0],
			Longitude: c[1],
			Timestamp: time.Now(),
		}

		fmt.Printf("[GPS] Train %s position: %.3f, %.3f\n", train.ID, c[0], c[1])
		gpsCh <- data
		time.Sleep(2 * time.Second)
	}
}
