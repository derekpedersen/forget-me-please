package twitter

import (
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
	opt["P"] = model.Option{
		Key:     "P",
		Value:   "P",
		Display: "(P)urge Twitter",
		Action:  PurgeTwitter,
	}
	log.WithField("TwitterOptions", opt).Debug("NewTwitterOptions")
	return opt
}

func Unlike() error {
	twts, err := NewTweetsLiked(auth, user, nil)
	if err != nil {
		log.Fatal(err)
	}
	return Paginate(twts, twts.Unlike, NewTweetsLiked)
}

func UnRetweet() error {
	twts, err := NewTweets(auth, user, nil)
	if err != nil {
		log.Fatal(err)
	}
	return Paginate(twts, twts.UnRetweet, NewTweets)
}

func DeleteTweets() error {
	twts, err := NewTweets(auth, user, nil)
	if err != nil {
		log.Fatal(err)
	}
	return Paginate(twts, twts.Delete, NewTweets)
}

func PurgeTwitter() error {
	Unlike()
	DeleteTweets()
	return nil
}

func Paginate(twts Tweets, action func() error, update func(auth Auth, user User, token *string) (Tweets, error)) error {
	for len(twts.Meta.NextToken) > 0 {
		twts, err = update(auth, user, &twts.Meta.NextToken)
		_ = action()
	}
	return nil
}
