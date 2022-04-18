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
		Action:       Unlike,
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

func Unlike() error {
	log.WithField("Unlike", time.Now())
	twts, err := NewTweetsLiked(auth, user, nil)
	if err != nil {
		log.Fatal(err)
	}
	return Paginate(twts, twts.Unlike, NewTweetsLiked)
}

func UnRetweet() error {
	log.WithField("UnRetweet", time.Now())
	twts, err := NewTweets(auth, user, nil)
	if err != nil {
		log.Fatal(err)
	}
	return Paginate(twts, twts.UnRetweet, NewTweets)
}

func DeleteTweets() error {
	log.WithField("DeleteTweet", time.Now())
	twts, err := NewTweets(auth, user, nil)
	if err != nil {
		log.Fatal(err)
	}
	return Paginate(twts, twts.Delete, NewTweets)
}

func PurgeTwitter() error {
	log.WithField("Purge", time.Now())
	Unlike()
	DeleteTweets()
	return nil
}

func Paginate(twts Tweets, action func() error, update func(auth Auth, user User, token *string) (Tweets, error)) error {
	for len(twts.Meta.NextToken) > 0 {
		// TODO: LOG METHOD AND TOKEN AND SUCH
		_ = action()
		twts, err = update(auth, user, &twts.Meta.NextToken)
		if err != nil {
			return err
		}
	}
	return nil
}
