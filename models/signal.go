package models

type SignalState string

const (
	Green  SignalState = "GREEN"
	Yellow SignalState = "YELLOW"
	Red    SignalState = "RED"
)

type Signal struct {
	ID      string
	Segment *Segment
	State   SignalState
}
