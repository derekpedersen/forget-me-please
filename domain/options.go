package domain

import (
	"fmt"

	"github.com/derekpedersen/forget-me-please/model"
)

type Options map[string]model.Option

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
