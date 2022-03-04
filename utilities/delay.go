package utilities

import (
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"
)

func Delay() {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(10) // n will be between 0 and 10
	log.Debugf("Sleeping %d seconds...\n", n)
	time.Sleep(time.Duration(n) * time.Second)
}
