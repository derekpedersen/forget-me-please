package twitter

import (
	"github.com/derekpedersen/forget-me-please/utilities"
	log "github.com/sirupsen/logrus"
)

type Tweets struct {
	Config Config
	User   User
	Data   []Tweet `json:"data"`
	Meta   struct {
		OldestId    string `json:"oldest_id"`
		NewestId    string `json:"newest_id"`
		ResultCount int    `json:"result_count"`
		NextToken   string `json:"next_token"`
	}
}

func NewTweets(config Config, user User, paginationToken *string, likedTweets *bool) (Tweets, error) {
	if len(config.Archive) > 0 {
		return newArchivedTweets(config, likedTweets)
	}
	return newTimeLineTweets(config, user, paginationToken, likedTweets)
}

func (twts *Tweets) Unlike() error {
	for _, v := range twts.Data {
		utilities.Delay()
		if !v.IsExempt(twts.Config.TwitterExemptUsers) {
			err := v.Unlike(twts.Config, twts.User)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (twts *Tweets) Delete() error {
	for _, v := range twts.Data {
		utilities.Delay()
		if !v.IsExempt(twts.Config.TwitterExemptUsers) {
			err := v.Delete(twts.Config, twts.User)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (twts *Tweets) UnRetweet() error {
	for _, v := range twts.Data {
		utilities.Delay()
		if v.IsRetweet() && !v.IsExempt(twts.Config.TwitterExemptUsers) {
			err := v.Delete(twts.Config, twts.User)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	return nil
}
