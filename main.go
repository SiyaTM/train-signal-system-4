// Now that we’re introducing the signal system, we can make the GPS the central broadcaster for movement updates —
// sensors listen to those updates and trigger signal changes accordingly.
package main

import (
	"signal-system4/models"
	"signal-system4/simulation"
	"time"
)

func main() {
	segments := []*models.Segment{
		{Name: "Segment-A", Latitude: -25.755, Longitude: 28.231},
		{Name: "Segment-B", Latitude: -25.756, Longitude: 28.232},
		{Name: "Segment-C", Latitude: -25.757, Longitude: 28.233},
	}

	trains := []*models.Train{{ID: "Train-1"}}

	gpsCh := make(chan *models.GPSData)
	signalCh := make(chan *models.SensorEvent)

	signalSystem := simulation.NewSignalSystem(segments)
	sensorSystem := simulation.NewSensorSystem(segments)

	go signalSystem.ListenForSensorEvents(signalCh)
	go sensorSystem.ListenForGPS(gpsCh, signalCh)

	for _, train := range trains {
		go simulation.StartGPSUpdates(train, gpsCh)
	}

	time.Sleep(20 * time.Second)
}
