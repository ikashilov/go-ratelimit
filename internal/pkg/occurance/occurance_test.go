package occurance

import (
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestOccurance(t *testing.T) {
	rate := New(0.1)

	for i := 0; i < 1000; i++ {
		// 1000 events in 1 second
		time.Sleep(1 * time.Millisecond)
		rate.Update()
	}

	// 10% error
	log.Println(rate.Get())
	assert.InDelta(t, rate.Get(), 100., 10.)
}
