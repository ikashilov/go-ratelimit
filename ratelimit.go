package ratelimit

import (
	"fmt"

	"github.com/ikashilov/go-ratelimit/internal/pkg/occurance"
)

// RateLimit holds an estimated rolling speed of events
// and a maximum speed for such events allowed to happen
// allowedRate is event per second
type RateLimit struct {
	allowedEventOccurance float64
	occuranceEstimator    *occurance.OccurrenceRate
}

// New creates a new Rate
func New(maxOccurance int, smoothing float64) (*RateLimit, error) {
	if smoothing < 0 || smoothing > 1 {
		return nil, fmt.Errorf("smoothing value should be in range (0, 1)")
	}

	if maxOccurance <= 0 {
		return nil, fmt.Errorf("max occurance should be greater than 0")
	}

	return &RateLimit{
		allowedEventOccurance: float64(maxOccurance),
		occuranceEstimator:    occurance.New(smoothing),
	}, nil
}

// Allow checks whether the event is allowed to happen
func (r *RateLimit) Allow() (bool, float64) {
	curOccurance := r.occuranceEstimator.Update()
	return curOccurance < r.allowedEventOccurance, curOccurance
}
