package domain

import (
	"time"

	"github.com/derekpedersen/forget-me-please/model"
	log "github.com/sirupsen/logrus"
)

type Option model.Option

func ProcessOption(opt *Option) {
	log.WithField("ProcessOptions Runtime", time.Now())
	if opt != nil && opt.Action != nil {
		log.Printf(opt.Confirmation)
		opt.Action()
		return
	}
	log.Println("Not yet supported sorry")
}
