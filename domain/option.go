package domain

import (
	"fmt"
	"log"

	"github.com/derekpedersen/forget-me-please/model"
)

type Option model.Option

func ProcessOption(opt *Option) {
	if opt == nil {
		fmt.Println("Not yet supported sorry")
	} else {
		log.Printf(opt.Confirmation)
		opt.Action()
	}
}
