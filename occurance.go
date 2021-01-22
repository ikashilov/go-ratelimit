package main

import (
	"math"
	"time"
)

// =============================================
// Based on https://stackoverflow.com/a/23617678
// =============================================

// DeafultSmoothing is a bearshit
const DeafultSmoothing = 0.1

// OccurrenceRate is a bullshit
type OccurrenceRate struct {
	uLast float64
	tLast float64
	k     float64
}

// NewOccurrenceRate is cowshit
func NewOccurrenceRate(k float64) OccurrenceRate {
	return OccurrenceRate{
		uLast: 0,
		tLast: float64(time.Now().UnixNano()) / 1000 / 1000 / 1000,
		k:     k,
	}
}

// Update is a pigshit
func (s *OccurrenceRate) Update() float64 {
	tNow := float64(time.Now().UnixNano()) / 1000.0 / 1000.0 / 1000.0

	s.uLast = math.Exp(s.k*(s.tLast-tNow))*s.uLast + s.k
	s.tLast = tNow

	return s.uLast
}

// GetLastSpeed is a dogshit
func (s *OccurrenceRate) GetLastSpeed() float64 {
	return s.uLast
}
