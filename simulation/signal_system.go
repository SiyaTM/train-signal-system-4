// We’ll use the 3-state signal logic you already approved (Green → Yellow → Red).
package simulation

import (
	"fmt"
	"signal-system4/models"
	"sync"
)

type SignalSystem struct {
	Signals []*models.Signal
	mu      sync.Mutex
}

func NewSignalSystem(segments []*models.Segment) *SignalSystem {
	signals := make([]*models.Signal, len(segments))

	for i, seg := range segments {
		signals[i] = &models.Signal{
			ID:      fmt.Sprintf("Signal-%d", i),
			Segment: seg,
			State:   models.Green,
		}
	}

	return &SignalSystem{Signals: signals}
}

func (ss *SignalSystem) UpdateSignal(signal *models.Signal, state models.SignalState) {
	ss.mu.Lock()
	defer ss.mu.Unlock()
	signal.State = state
	fmt.Printf("[SIGNAL] %s for %s → %s\n", signal.ID, signal.Segment.Name, state)
}

func (ss *SignalSystem) ListenForSensorEvents(signalCh <-chan *models.SensorEvent) {
	for ev := range signalCh {
		for i, signal := range ss.Signals {
			if signal.Segment == ev.Segment {
				switch ev.Type {
				case models.TrainEntered:
					ss.UpdateSignal(signal, models.Red)
					if i > 0 {
						prev := ss.Signals[i-1]
						ss.UpdateSignal(prev, models.Yellow)
					}
				case models.TrainExited:
					ss.UpdateSignal(signal, models.Green)
					if i > 0 {
						prev := ss.Signals[i-1]
						ss.UpdateSignal(prev, models.Green)
					}
				}
			}
		}
	}
}
