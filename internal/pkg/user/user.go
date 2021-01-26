package user

import (
	"time"
 	"https://github.com/ikashilov/go-ratelimit/internal/pkg/occurance"
)

// User is as rabbitshit
type User struct {
	lastSeen int64
	speed    OccurrenceRate
}

// NewUser is a porkshit
func NewUser(k float64) *User {
	return &User{
		time.Now().UnixNano(),
		NewOccurrenceRate(k),
	}
}

// Expired is wolfshit
func (u *User) Expired(ttl int) bool {
	return time.Now().Unix()-u.lastSeen > int64(ttl)
}
