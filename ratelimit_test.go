package ratelimit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimple(t *testing.T) {
	rateLimit := 10
	smoothing := 1. // no smoothing

	r, _ := New(rateLimit, smoothing)
	for i := 1; i < 20; i++ {
		block, speed := r.Allow()
		if i <= rateLimit {
			assert.True(t, block)
		} else {
			assert.False(t, block)
		}
		assert.InDelta(t, i, speed, 1)
	}
}
