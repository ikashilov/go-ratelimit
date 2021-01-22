package main

import (
	"log"
	"time"
)

func main() {
	s := NewUserBucket(5, 60, DeafultCleanUP, DeafultSmoothing)
	s.Start()

	userID := "hooy"
	waitTime := 1000 * time.Millisecond

	for i := 0; i < 1000; i++ {
		time.Sleep(50 * time.Millisecond)
		speed, allow := s.Allow(userID)
		if allow {
			log.Printf("Allow --> %.3f\n", speed)
		} else {
			log.Printf("Block --> %.3f. Sleeping for %v\n", speed, waitTime)
			time.Sleep(waitTime)
		}
	}
}
