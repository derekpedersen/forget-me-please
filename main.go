package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/derekpedersen/forget-me-please/domain"
	"github.com/derekpedersen/forget-me-please/model"
	"github.com/derekpedersen/forget-me-please/twitter"
	"github.com/derekpedersen/forget-me-please/utilities"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithField("StartTime", time.Now())
	log.WithField("LogLevel", log.GetLevel())
	log.Info("**** FORGET-ME-PLEASE ****")
	flag.Parse()

	// TODO: need an option to go with a CLI or self hosted PWA

	socialMedia := NewSocialMediaOptions()
	socialMedia.PrintOptions()
	key := utilities.ReadLine()
	opt := socialMedia.SelectOption(*key)

	if opt == nil {
		fmt.Println("Not yet supported sorry")
	} else {
		opt.Action()
	}
}

func NewSocialMediaOptions() domain.Options {
	opt := domain.Options{}
	// TODO: this should be driven by a db or a least a json file
	opt["T"] = model.Option{
		Key:     "T",
		Value:   "Twitter",
		Display: "(T)witter",
		Action:  twitter.Twitter,
	}
	log.WithField("SocialMedia", opt).Debug("NewSocialMediaOptions")
	return opt
}
