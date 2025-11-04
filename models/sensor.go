package models

type SensorEventType string

const (
	TrainEntered SensorEventType = "ENTERED"
	TrainExited  SensorEventType = "EXITED"
)

type Sensor struct {
	ID      string
	Segment *Segment
}

type SensorEvent struct {
	SensorID string
	TrainID  string
	Type     SensorEventType
	Segment  *Segment
}
