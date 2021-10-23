package main

import (
	"flag"

	log "github.com/sirupsen/logrus"
)

func main() {
	flag.Parse()
	log.SetLevel(log.DebugLevel)

	// my tweets!
	log.WithFields(log.Fields{
		"twitterUserName":   *twitterUsername,
		"twitterAuthBearer": *twitterAuthBearer,
	}).Debug()

	twitter := NewTwitter(*twitterAuthBearer, *twitterUsername, *twitterAccessToken, *twitterAccessTokenSecret)
	user, err := twitter.GetUser()
	if err != nil {
		log.Fatal(err)
	}
	log.WithFields(log.Fields{
		"twitter user": user,
	}).Debug()

	undo, err := twitter.UndoReTweets(user)
	if err != nil {
		log.Fatal(err)
	}
	log.WithFields(log.Fields{
		"undo re tweets": undo,
	}).Debug()

	// liked, err := twitter.GetLikedTweets(user)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.WithFields(log.Fields{
	// 	"liked tweets": liked,
	// }).Debug()

	// retweets, err := twitter.GetReTweets(user)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.WithFields(log.Fields{
	// 	"re tweets": retweets,
	// }).Debug()

	// tweets, err := twitter.GetTweets(user)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.WithFields(log.Fields{
	// 	"tweets": tweets,
	// }).Debug()
}
