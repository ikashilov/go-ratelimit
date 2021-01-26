package occurance

import (
	"math"
	"time"
)

// =============================================
// Based on https://stackoverflow.com/a/23617678
// =============================================

// DeafultSmoothing is a default `smoothing` paramater to estimate occurance
const DeafultSmoothing = 0.1

// OccurrenceRate is a exponential event-occurance estimator
type OccurrenceRate struct {
	uLast float64
	tLast float64
	k     float64
}

// New creates an OccurrenceRate with a given `smoothing parameter`
func New(k float64) *OccurrenceRate {
	return &OccurrenceRate{
		uLast: 0,
		tLast: float64(time.Now().UnixNano()) / 1000 / 1000 / 1000,
		k:     k,
	}
}

// Update update the occurance rate and returns the last value of the estimted occurance
func (s *OccurrenceRate) Update() float64 {
	tNow := float64(time.Now().UnixNano()) / 1000.0 / 1000.0 / 1000.0

	s.uLast = math.Exp(s.k*(s.tLast-tNow))*s.uLast + s.k
	s.tLast = tNow

	return s.uLast
}

// Get returns the last value of the estimted occurnace
// which can be read as `n events per second`
func (s *OccurrenceRate) Get() float64 {
	return s.uLast
}
