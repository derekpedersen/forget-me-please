package domain

import (
	"fmt"
	"log"

	"github.com/derekpedersen/forget-me-please/model"
)

type Option model.Option

func ProcessOption(opt *Option) {
	if opt != nil && opt.Action != nil {
		log.Printf(opt.Confirmation)
		opt.Action()
		return
	}
	fmt.Println("Not yet supported sorry")
}
