package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/ikashilov/go-ratelimit"
	"github.com/ikashilov/go-ratelimit/internal/pkg/occurance"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	s := ratelimit.NewUserBucket(5, 60, ratelimit.DeafultCleanUP, occurance.NoSmoothing)
	s.Start()

	userID := "this can be a user id" // "127.0.0.1"

	for i := 0; i < 100; i++ {
		time.Sleep(100 * time.Millisecond)
		speed, allow := s.Allow(userID)
		if allow {
			log.Printf("Allow --> %.3f\n", speed)
		} else {
			waitTime := time.Duration(rand.Intn(5)) * time.Second
			log.Printf("Block --> %.3f. Sleeping for %v\n", speed, waitTime)
			time.Sleep(waitTime)
		}
	}
	// rateLimit := 5 // 5 req/sec
	// r, _ := ratelimit.New(rateLimit, occurance.NoSmoothing)
	// for i := 1; i < 10; i++ {
	// 	block, speed := r.Allow()
	// 	fmt.Println("requests #", i, "allow:", block, "speed:", speed)
	// }
}
