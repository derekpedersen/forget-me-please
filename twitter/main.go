package twitter

import (
	"fmt"

	"github.com/derekpedersen/forget-me-please/utilities"
	log "github.com/sirupsen/logrus"
)

var auth Auth
var user User
var err error

func Twitter() error {
	log.Debug("Twitter")
	auth = NewAuth()
	user, err = NewUser(auth)
	if err != nil {
		log.Fatal(err)
	}
	log.WithFields(log.Fields{
		"twitter user": user,
	}).Debug()

	opts := NewOptions()
	opts.PrintOptions()
	key := utilities.ReadLine(utilities.Reader())
	opt := opts.SelectOption(*key)
	if opt == nil {
		fmt.Println("Not yet supported sorry")
	} else {
		opt.Action()

	}
	log.Debugf("Completed Option: %v", opt.Display)
	return nil
}
