package twitter

import (
	"math/rand"
	"time"

	"github.com/derekpedersen/forget-me-please/domain"
	"github.com/derekpedersen/forget-me-please/model"
	log "github.com/sirupsen/logrus"
)

func NewOptions() domain.Options {
	opt := domain.Options{}
	// TODO: this should be driven by a db or a least a json file
	opt["L"] = model.Option{
		Key:     "L",
		Value:   "Unlike",
		Display: "Un(L)ike",
		Action:  Unlike,
	}
	opt["R"] = model.Option{
		Key:     "R",
		Value:   "Unretweet",
		Display: "Un(R)etweet",
		Action:  UnRetweet,
	}
	opt["D"] = model.Option{
		Key:     "D",
		Value:   "Delete",
		Display: "(D)elete Tweets",
		Action:  DeleteTweets,
	}
	// opt["P"] = model.Option{
	// 	Key:     "P",
	// 	Value:   "P",
	// 	Display: "(P)urge Twitter",
	// 	Action:  PurgeTwitter,
	// }
	log.WithField("TwitterOptions", opt).Debug("NewTwitterOptions")
	return opt
}

func Unlike() error {
	newTweets, err := NewTweetsLiked(auth, user, nil)
	if err != nil {
		log.Fatal(err)
	}
	_ = newTweets.Unlike()
	for len(newTweets.Meta.NextToken) > 0 {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(10) // n will be between 0 and 10
		log.Debugf("Sleeping %d seconds...\n", n)
		time.Sleep(time.Duration(n) * time.Second)
		newTweets, err = NewTweets(auth, user, &newTweets.Meta.NextToken)
		if err != nil {
			log.Fatal(err)
		}
		_ = newTweets.Unlike()
	}

	return nil
}

func UnRetweet() error {
	newTweets, err := NewTweets(auth, user, nil)
	if err != nil {
		log.Fatal(err)
	}
	_ = newTweets.UnRetweet()
	for len(newTweets.Meta.NextToken) > 0 {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(10) // n will be between 0 and 10
		log.Debugf("Sleeping %d seconds...\n", n)
		time.Sleep(time.Duration(n) * time.Second)
		newTweets, err = NewTweets(auth, user, &newTweets.Meta.NextToken)
		if err != nil {
			log.Fatal(err)
		}
		_ = newTweets.UnRetweet()
	}
	return nil
}

func DeleteTweets() error {
	newTweets, err := NewTweets(auth, user, nil)
	if err != nil {
		log.Fatal(err)
	}
	_ = newTweets.Delete()
	for len(newTweets.Meta.NextToken) > 0 {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(10) // n will be between 0 and 10
		log.Debugf("Sleeping %d seconds...\n", n)
		time.Sleep(time.Duration(n) * time.Second)
		newTweets, err = NewTweets(auth, user, &newTweets.Meta.NextToken)
		if err != nil {
			log.Fatal(err)
		}
		_ = newTweets.Delete()
	}
	return nil
}

func PurgeTwitter() error {
	Unlike()
	DeleteTweets()
	return nil
}
