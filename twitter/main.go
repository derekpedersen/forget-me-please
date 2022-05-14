package twitter

import (
	"github.com/derekpedersen/forget-me-please/domain"
	"github.com/derekpedersen/forget-me-please/utilities"
	log "github.com/sirupsen/logrus"
)

var auth Auth
var user User
var err error

func Twitter() error {
	log.Printf("Twitter")
	auth = NewAuth()
	user, err = NewUser(auth)
	if err != nil {
		log.Fatal(err)
	}
	log.WithFields(log.Fields{
		"twitter user": user,
	}).Debug("Twitter")

	opts := NewOptions()
	opts.PrintOptions()
	key := utilities.ReadLine(utilities.Reader())
	opt := opts.SelectOption(*key)
	domain.ProcessOption(opt)
	log.Printf("Completed Option: %v", opt.Display)
	return nil
}
