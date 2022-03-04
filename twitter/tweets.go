package twitter

import (
	"encoding/json"
	"net/http"

	"github.com/derekpedersen/forget-me-please/utilities"
	log "github.com/sirupsen/logrus"
)

// {"data":[{"id":"1426928113426993152","text":"All of you guys are making fun of Marianne Williamson because you don't have any better ideas. Let's give it a shot. Let's deploy Jimmy Dore to Afghanistan."}]
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

func NewTweets(auth Auth, user User, paginationToken *string) (Tweets, error) {
	var tweets Tweets
	url := "https://api.twitter.com/2/users/" + user.Data.ID + "/tweets?max_results=100"
	if paginationToken != nil && len(*paginationToken) > 0 {
		url += "&pagination_token=" + *paginationToken
	}
	data, err := utilities.HttpRequest(url, http.MethodGet, auth.AuthorizationBearerToken())
	if err != nil {
		log.Errorf("Error performing request:\n %v", err)
		return tweets, err
	}
	log.Debugf("GetTweets: %v", data)

	if err = json.Unmarshal([]byte(*data), &tweets); err != nil {
		log.Error(err)
		return tweets, err
	}
	tweets.Auth = auth
	tweets.User = user
	return tweets, nil
}

func NewTweetsLiked(auth Auth, user User, paginationToken *string) (Tweets, error) {
	var tweets Tweets
	url := "https://api.twitter.com/2/users/" + user.Data.ID + "/liked_tweets?max_results=100"
	if paginationToken != nil && len(*paginationToken) > 0 {
		url += "&pagination_token=" + *paginationToken
	}
	data, err := utilities.HttpRequest(url, http.MethodGet, auth.AuthorizationBearerToken())
	if err != nil {
		log.Errorf("Error performing request:\n %v", err)
		return tweets, err
	}
	log.Debugf("GetLikedTweets: %s", data)

	if err = json.Unmarshal([]byte(*data), &tweets); err != nil {
		log.Error(err)
		return tweets, err
	}

	return tweets, nil
}

func (twts *Tweets) Unlike() error {
	for _, v := range twts.Data {
		err := v.Unlike(twts.Auth, twts.User)
		if err != nil {
			return err
		}
	}
	return nil
}

func (twts *Tweets) Delete() error {
	for _, v := range twts.Data {
		err := v.Delete(twts.Auth, twts.User)
		if err != nil {
			return err
		}
	}
	return nil
}

func (twts *Tweets) UnRetweet() error {
	for _, v := range twts.Data {
		if v.IsRetweet() && !v.IsExempt(twts.Auth.TwitterExemptUsers) {
			err := v.Delete(twts.Auth, twts.User)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	return nil
}
