package main

import (
	"fmt"
	"time"

	"github.com/derekpedersen/forget-me-please/twitter"
	"github.com/derekpedersen/forget-me-please/utilities"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithField("StartTime", time.Now())
	log.WithField("LogLevel", log.GetLevel())
	log.Info("**** FORGET-ME-PLEASE ****")

	// TODO: need an option to go with a CLI or self hosted PWA

	fmt.Println("(F)acebook")
	fmt.Println("(R)eddit")
	fmt.Println("(T)witter")
	fmt.Println("(Y)ouTube")
	fmt.Println("(I)nstagram")

	text := utilities.ReadLine()
	log.WithField("SocialMedia", *text).Debug()

	switch *text {
	case "T":
		twitter.TakeAction()
	default:
		fmt.Println("Not yet supported sorry")
	}
}
