package domain

import (
	"fmt"

	"github.com/derekpedersen/forget-me-please/model"
	"github.com/derekpedersen/forget-me-please/twitter"
	log "github.com/sirupsen/logrus"
)

type Options map[string]model.Option

func NewSocialMediaOptions() Options {
	opt := Options{}
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

func NewTwitterOptions() Options {
	opt := Options{}

	return opt
}

func (dom Options) PrintOptions() {
	for _, v := range dom {
		fmt.Println(v.Display)
	}
}

func (dom Options) SelectOption(key string) *Option {
	for k, v := range dom {
		if k == key {
			opt := Option(v)
			return &opt
		}
	}

	return nil
}
