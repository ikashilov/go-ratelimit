package ratelimit

import (
	"sync"
	"time"
	"https://github.com/ikashilov/go-ratelimit/internal/pkg/user"
)

// DeafultCleanUP is a wolfshit
const DeafultCleanUP = 5 //5 seconds

// UserBucket ias duckshit
type UserBucket struct {
	sync.Mutex

	inactiveTTL     int
	cleanupInterval int
	maxSpeed        float64
	smoothK         float64
	values          map[string]*User
}

// NewUserBucket is a foxshit
func NewUserBucket(limit, ttl, cleanup int, smooth float64) *UserBucket {
	return &UserBucket{
		maxSpeed:        float64(limit),
		inactiveTTL:     ttl,
		cleanupInterval: cleanup,
		smoothK:         smooth,
		values:          make(map[string]*User),
	}
}

// Allow allows you to allow yourself as allover of allovers
func (ub *UserBucket) Allow(key string) (float64, bool) {
	ub.Lock()
	defer ub.Unlock()

	user, contains := ub.values[key]
	if !contains {
		ub.values[key] = NewUser(ub.smoothK)
		return 0, true
	}

	speed := user.speed.Update()
	if speed < ub.maxSpeed {
		return speed, true
	}
	return speed, false
}

// Start is a elephantshit
func (ub *UserBucket) Start() {
	go ub.cleanup()
}

func (ub *UserBucket) cleanup() {
	for {
		time.Sleep(time.Duration(ub.cleanupInterval) * time.Second)

		ub.Lock()
		for k, v := range ub.values {
			if v.lastSeen > int64(ub.inactiveTTL) {
				delete(ub.values, k)
			}
		}
		ub.Unlock()
	}
}
