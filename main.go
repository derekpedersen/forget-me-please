package main

import (
	"flag"

	"github.com/derekpedersen/forget-me-please/twitter"
	log "github.com/sirupsen/logrus"
)

func main() {
	flag.Parse()
	log.SetLevel(log.DebugLevel)
	auth := twitter.NewTwitterAuth()
	user, err := twitter.NewTwitterUser(auth)
	if err != nil {
		log.Fatal(err)
	}
	log.WithFields(log.Fields{
		"twitter user": user,
	}).Debug()

	newTweets, err := twitter.NewTweets(auth, user)
	if err != nil {
		log.Fatal(err)
	}
	_ = newTweets.UnRetweet()
}
