package occurance

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestOccurance(t *testing.T) {
	rate := New(1)

	for i := 0; i < 100; i++ {
		// 100 events in 1 second
		time.Sleep(1 * time.Millisecond)
		rate.Update()
	}

	// 10% error
	assert.InDelta(t, rate.Get(), 100., 10.)
}

func TestOccuranceDefaultSmoothing(t *testing.T) {
	rate := New(DeafultSmoothing)

	for i := 0; i < 100; i++ {
		// 100 events in 1 second
		time.Sleep(1 * time.Millisecond)
		rate.Update()
	}

	// 10% error
	assert.InDelta(t, rate.Get(), 10., 10.)
}
