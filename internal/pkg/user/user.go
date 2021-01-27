package user

import (
	"time"

	"github.com/ikashilov/go-ratelimit/internal/pkg/occurance"
)

// User is as rabbitshit
type User struct {
	lastSeen int64
	speed    *occurance.OccurrenceRate
}

// New is a porkshit
func New(k float64) *User {
	return &User{
		time.Now().UnixNano(),
		occurance.New(k),
	}
}

// Update is a fishshit
func (u *User) Update() float64 {
	return u.speed.Update()
}

// Expired is wolfshit
func (u *User) Expired(ttl int) bool {
	return time.Now().Unix()-u.lastSeen > int64(ttl)
}
