package twitter

import (
	"time"

	"github.com/derekpedersen/forget-me-please/domain"
	"github.com/derekpedersen/forget-me-please/model"
	log "github.com/sirupsen/logrus"
)

func NewOptions() domain.Options {
	opt := domain.Options{}
	opt["L"] = model.Option{
		Key:          "L",
		Value:        "Unlike",
		Display:      "Un(L)ike",
		Action:       UnlikeTweets,
		Confirmation: "Proceeding to Unlike Tweets",
	}
	opt["R"] = model.Option{
		Key:          "R",
		Value:        "Unretweet",
		Display:      "Un(R)etweet",
		Action:       UnRetweet,
		Confirmation: "Proceeding to Unretweet Tweets",
	}
	opt["D"] = model.Option{
		Key:          "D",
		Value:        "Delete",
		Display:      "(D)elete Tweets",
		Action:       DeleteTweets,
		Confirmation: "Proceeding to Delete Tweets",
	}
	opt["P"] = model.Option{
		Key:          "P",
		Value:        "P",
		Display:      "(P)urge Twitter",
		Action:       PurgeTwitter,
		Confirmation: "Proceeding to Purge Twitter",
	}
	log.WithField("TwitterOptions", opt).Debug("NewTwitterOptions")
	return opt
}

func UnlikeTweets() error {
	likedtweets := true
	log.WithField("Unlike", time.Now())
	twts, err := NewTweets(config, user, nil, &likedtweets)
	if err != nil {
		log.Fatal(err)
		return err
	}
	twts.Unlike()
	for len(twts.Meta.NextToken) > 0 {
		twts, err = NewTweets(config, user, &twts.Meta.NextToken, &likedtweets)
		if err != nil {
			log.Fatal(err)
			return err
		}
		twts.Unlike()
	}
	return nil
}

func UnRetweet() error {
	log.WithField("UnRetweet", time.Now())
	twts, err := NewTweets(config, user, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	twts.UnRetweet()
	for len(twts.Meta.NextToken) > 0 {
		twts, err = NewTweets(config, user, &twts.Meta.NextToken, nil)
		if err != nil {
			log.Fatal(err)
			return err
		}
		twts.UnRetweet()
	}
	return nil
}

func DeleteTweets() error {
	log.WithField("DeleteTweet", time.Now())
	twts, err := NewTweets(config, user, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	twts.Delete()
	for len(twts.Meta.NextToken) > 0 {
		twts, err = NewTweets(config, user, &twts.Meta.NextToken, nil)
		if err != nil {
			log.Fatal(err)
			return err
		}
		twts.Delete()
	}
	return nil
}

func PurgeTwitter() error {
	log.WithField("Purge", time.Now())
	UnlikeTweets()
	DeleteTweets()
	return nil
}
