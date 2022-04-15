package main

import (
	"flag"
	"time"

	"github.com/derekpedersen/forget-me-please/domain"
	"github.com/derekpedersen/forget-me-please/model"
	"github.com/derekpedersen/forget-me-please/twitter"
	"github.com/derekpedersen/forget-me-please/utilities"
	log "github.com/sirupsen/logrus"
)

func main() {
	// configure logrus
	log.WithField("StartTime", time.Now())
	log.WithField("LogLevel", log.GetLevel())
	log.Info("**** FORGET-ME-PLEASE ****")

	// flags (command line args)
	flag.Parse()

	// TODO: need an option to go with a CLI or self hosted PWA

	// handle which social media we are interacting with
	socialMedia := NewSocialMediaOptions()
	socialMedia.PrintOptions()
	key := utilities.ReadLine(utilities.Reader())
	opt := socialMedia.SelectOption(*key)
	domain.ProcessOption(opt)
}

func NewSocialMediaOptions() domain.Options {
	opt := domain.Options{}
	// TODO: more options than just twitter
	opt["T"] = model.Option{
		Key:          "T",
		Value:        "Twitter",
		Display:      "(T)witter",
		Action:       twitter.Twitter,
		Confirmation: "Lets Forget About Some Tweets",
	}
	log.WithField("SocialMedia", opt).Debug("NewSocialMediaOptions")
	return opt
}
