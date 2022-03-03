package twitter

import (
	log "github.com/sirupsen/logrus"
)

func Twitter() {
	log.Info("Twitter")
	auth := NewTwitterAuth()
	user, err := NewTwitterUser(auth)
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
	_ = newTweets.UnRetweet()
	for len(newTweets.Meta.NextToken) > 0 {
		newTweets, err = NewTweets(auth, user, &newTweets.Meta.NextToken)
		if err != nil {
			log.Fatal(err)
		}
		_ = newTweets.UnRetweet()
	}
}
