package ratelimit

import (
	"sync"
	"time"

	"github.com/ikashilov/go-ratelimit/internal/pkg/user"
)

// DeafultCleanUP is a default period for GC to step out the scene
const DeafultCleanUP = 1 //1 Minute

// UserBucket is a thread safe way to manage requests with ratelimt
type UserBucket struct {
	sync.Mutex

	inactiveTTL     int
	cleanupInterval int
	maxSpeed        float64
	smoothK         float64
	values          map[string]*user.User
}

// NewUserBucket creates a new user bucket
func NewUserBucket(limit, ttl, cleanup int, smooth float64) *UserBucket {
	return &UserBucket{
		maxSpeed:        float64(limit),
		inactiveTTL:     ttl,
		cleanupInterval: cleanup,
		smoothK:         smooth,
		values:          make(map[string]*user.User),
	}
}

// Allow checks whether a user with given `key` is allowed to procced
// and return the estimated speed for such user
func (ub *UserBucket) Allow(key string) (float64, bool) {
	ub.Lock()
	defer ub.Unlock()

	val, contains := ub.values[key]
	if !contains {
		ub.values[key] = user.New(ub.smoothK)
		return 0, true
	}

	speed := val.Update()
	if speed < ub.maxSpeed {
		return speed, true
	}
	return speed, false
}

// Start runs a grabge collenctor for not used users
func (ub *UserBucket) Start() {
	go ub.cleanup()
}

// Garbage collector: evicts uused users
func (ub *UserBucket) cleanup() {
	for {
		time.Sleep(time.Duration(ub.cleanupInterval) * time.Second)

		ub.Lock()
		for k, v := range ub.values {
			if v.Expired(ub.inactiveTTL) {
				delete(ub.values, k)
			}
		}
		ub.Unlock()
	}
}
