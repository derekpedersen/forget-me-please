package twitter

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/derekpedersen/forget-me-please/utilities"
	log "github.com/sirupsen/logrus"
)

type Tweets struct {
	Auth Auth
	User User
	Data []Tweet `json:"data"`
	Meta struct {
		OldestId    string `json:"oldest_id"`
		NewestId    string `json:"newest_id"`
		ResultCount int    `json:"result_count"`
		NextToken   string `json:"next_token"`
	}
}

func NewTweets(auth Auth, user User, paginationToken *string, likedTweets *bool) (Tweets, error) {
	log.WithField("NewTweets", time.Now())
	var tweets Tweets
	url := "https://api.twitter.com/2/users/" + user.Data.ID
	if likedTweets != nil && *likedTweets {
		url += "/liked_tweets"
	} else {
		url += "/tweets"
	}
	url += "?max_results=100"
	// this line tries to avoid cached responses from twitter
	// url += "&" + utilities.Random() + "=" + utilities.Random()
	if paginationToken != nil && len(*paginationToken) > 0 {
		url += "&pagination_token=" + *paginationToken
	}
	data, err := utilities.HttpRequest(url, http.MethodGet, auth.AuthorizationBearerToken())
	if err != nil {
		log.Errorf("Error performing request:\n %v", err)
		return tweets, err
	}
	log.WithFields(log.Fields{"Tweets": data}).Debug("NewTweets")

	if err = json.Unmarshal([]byte(*data), &tweets); err != nil {
		log.Error(err)
		return tweets, err
	}
	tweets.Auth = auth
	tweets.User = user
	log.WithFields(log.Fields{"Tweets": tweets, "API Response": data, "URL": url}).Debug("NewTweets")
	return tweets, nil
}

func NewArchivedTweets() {}

func NewTimelineTweets() {}

func (twts *Tweets) Unlike() error {
	for _, v := range twts.Data {
		utilities.Delay()
		if !v.IsExempt(twts.Auth.TwitterExemptUsers) {
			err := v.Unlike(twts.Auth, twts.User)
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
		if !v.IsExempt(twts.Auth.TwitterExemptUsers) {
			err := v.Delete(twts.Auth, twts.User)
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
		if v.IsRetweet() && !v.IsExempt(twts.Auth.TwitterExemptUsers) {
			err := v.Delete(twts.Auth, twts.User)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	return nil
}
