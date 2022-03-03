package main

import (
	"flag"

	"github.com/derekpedersen/forget-me-please/twitter"
	log "github.com/sirupsen/logrus"
)

func main() {
	flag.Parse()
	log.SetLevel(log.DebugLevel)
	twitter.Twitter()
}
