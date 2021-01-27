package user

import (
	"time"

	"github.com/ikashilov/go-ratelimit/internal/pkg/occurance"
)

// User is a user struct to hold user estimated speed and `lastSeen` timestamp
type User struct {
	lastSeen int64
	speed    *occurance.OccurrenceRate
}

// New creates a new user
func New(k float64) *User {
	return &User{
		time.Now().Unix(),
		occurance.New(k),
	}
}

// Update updates user's speed
func (u *User) Update() float64 {
	u.lastSeen = time.Now().UnixNano()
	return u.speed.Update()
}

// Expired checks whether the user observed to long
func (u *User) Expired(ttl int) bool {
	return time.Now().Unix()-u.lastSeen > int64(ttl)
}
