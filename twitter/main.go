package twitter

import (
	"fmt"

	"github.com/derekpedersen/forget-me-please/utilities"
	log "github.com/sirupsen/logrus"
)

func Twitter() error {
	log.Info("Twitter")
	auth := NewAuth()
	user, err := NewUser(auth)
	if err != nil {
		log.Fatal(err)
	}
	log.WithFields(log.Fields{
		"twitter user": user,
	}).Debug()

	newTweets, err := NewTweets(auth, user, nil)
	if err != nil {
		log.Fatal(err)
	}

	opts := NewOptions(newTweets)
	opts.PrintOptions()
	key := utilities.ReadLine()
	opt := opts.SelectOption(*key)
	if opt == nil {
		fmt.Println("Not yet supported sorry")
	} else {
		opt.Action()
		for len(newTweets.Meta.NextToken) > 0 {
			newTweets, err = NewTweets(auth, user, &newTweets.Meta.NextToken)
			if err != nil {
				log.Fatal(err)
			}
			opt.Action()
		}
	}
	log.Info("Completed Option: %v", opt.Display)
	return nil
}
