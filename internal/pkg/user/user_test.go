package user

import (
	"testing"
	"time"

	"github.com/ikashilov/go-ratelimit/internal/pkg/occurance"
	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	u := New(occurance.DeafultSmoothing)

	time.Sleep(1 * time.Second)
	assert.False(t, u.Expired(1))

	time.Sleep(2 * time.Second)
	assert.True(t, u.Expired(1))
}
