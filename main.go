package main

import (
	"flag"

	log "github.com/sirupsen/logrus"
)

var twitterAuthBearer = flag.String("twitterAuthBearer", "", "Twitter Authorization Bearer Token")
var twitterUsername = flag.String("twitterUsername", "", "Twitter User Name")

func main() {
	flag.Parse()
	log.SetLevel(log.DebugLevel)

	// my tweets!
	log.WithFields(log.Fields{
		"twitterUserName":   *twitterUsername,
		"twitterAuthBearer": *twitterAuthBearer,
	}).Debug()

	twitter := NewTwitter(*twitterAuthBearer, *twitterUsername)
	user, err := twitter.GetUser()
	if err != nil {
		log.Fatal(err)
	}
	log.WithFields(log.Fields{
		"twitter user": user,
	}).Debug()

	liked, err := twitter.GetLikedTweets(user)
	if err != nil {
		log.Fatal(err)
	}
	log.WithFields(log.Fields{
		"liked tweets": liked,
	}).Debug()

	retweets, err := twitter.GetReTweets(user)
	if err != nil {
		log.Fatal(err)
	}
	log.WithFields(log.Fields{
		"re tweets": retweets,
	}).Debug()
}
