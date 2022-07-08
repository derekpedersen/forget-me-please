package domain

import (
	"fmt"

	"github.com/derekpedersen/forget-me-please/model"
	log "github.com/sirupsen/logrus"
)

type Options map[string]model.Option

func (dom Options) PrintOptions() {
	log.WithField("Options", dom).Printf("PrintOptions")
	fmt.Println()
	for _, v := range dom {
		fmt.Println(v.Display)
	}
}

func (dom Options) SelectOption(key string) *Option {
	log.WithFields(log.Fields{"Options": dom, "key": key}).Printf("SelectOption")
	for k, v := range dom {
		if k == key {
			opt := Option(v)
			return &opt
		}
	}
	return nil
}
