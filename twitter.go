package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

// {"data":{"id":"1684445455","name":"Derek Pedersen","username":"PedersenDerek"}}
type TwitterUser struct {
	Data struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		UserName string `json:"username"`
	} `json:"data"`
}

//{"data":[{"id":"1426928113426993152","text":"All of you guys are making fun of Marianne Williamson because you don't have any better ideas. Let's give it a shot. Let's deploy Jimmy Dore to Afghanistan."}]
type Tweets struct {
	Data []struct {
		ID   string `json:"id"`
		Text string `json:"text"`
	} `json:"data"`
}

// Twitter interface
type Twitter interface {
	GetUser() (TwitterUser, error)
	GetReplies(user TwitterUser) (tweets Tweets, err error)
	GetTweets(user TwitterUser) (tweets Tweets, err error)
	GetReTweets(user TwitterUser) (tweets Tweets, err error)
	GetLikedTweets(user TwitterUser) (tweets Tweets, err error)
	UnlikeTweets(user TwitterUser) (response interface{}, err error)
	DeleteTweets(user TwitterUser) (response interface{}, err error)
	UndoReTweets(user TwitterUser) (response interface{}, err error)
}

// TwitterImpl struct
type TwitterImpl struct {
	authToken string
	userName  string
}

// NewTwitter creates a new album service
func NewTwitter(authToken, userName string) Twitter {
	return &TwitterImpl{
		authToken: authToken,
		userName:  userName,
	}
}

func (svc *TwitterImpl) request(url, method string) (*string, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Errorf("Error creating request:\n %v", err)
		return nil, err
	}
	req.Header.Add("Accept", "application/json")

	// Create a Bearer string by appending string access token
	var bearer = "Bearer " + svc.authToken
	fmt.Print(bearer)

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Errorf("Error making request:\n %v", err)
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Errorf("Error reading res.Body:\n %v", err)
		return nil, err
	}
	s := string(body)
	log.WithFields(log.Fields{
		"body": s,
	}).Debug()
	return &s, nil
}

func (svc *TwitterImpl) GetUser() (TwitterUser, error) {
	var twitterUser TwitterUser

	url := "https://api.twitter.com/2/users/by/username/" + svc.userName
	data, err := svc.request(url, http.MethodGet)
	if err != nil {
		log.Errorf("Error performing request:\n %v", err)
		return twitterUser, err
	}

	if err = json.Unmarshal([]byte(*data), &twitterUser); err != nil {
		log.Error(err)
		return twitterUser, err
	}

	return twitterUser, nil
}

func (svc *TwitterImpl) GetTweets(user TwitterUser) (tweets Tweets, err error) {
	if err != nil {
		log.Errorf("Error getting user:\n %v", err)
		return tweets, err
	}

	url := "https://api.twitter.com/2/users/" + user.Data.ID + "/tweets"
	data, err := svc.request(url, http.MethodGet)
	if err != nil {
		log.Errorf("Error performing request:\n %v", err)
		return tweets, err
	}
	log.Debugf("GetReTweets: %v", data)

	if err = json.Unmarshal([]byte(*data), &tweets); err != nil {
		log.Error(err)
		return tweets, err
	}

	return tweets, nil
}

func (svc *TwitterImpl) GetLikedTweets(user TwitterUser) (tweets Tweets, err error) {
	if err != nil {
		log.Errorf("Error getting user:\n %v", err)
		return tweets, err
	}
	url := "https://api.twitter.com/2/users/" + user.Data.ID + "/liked_tweets"
	data, err := svc.request(url, http.MethodGet)
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

// curl -X DELETE https://api.twitter.com/2/users/:user-id/likes/:tweet-id -H "Authorization: OAuth $OAUTH_SIGNATURE"
// {
// 	"data": {
// 	  "liked": false
// 	}
// }
func (svc *TwitterImpl) UnlikeTweets(user TwitterUser) (response interface{}, err error) {
	liked, err := svc.GetLikedTweets(user)
	if err != nil {
		log.Fatal(err)
	}
	log.Debug(liked)

	//TODO: implement undo like
	for k, v := range liked.Data {
		log.Debug(k, v)
	}

	// TODO: update to real response
	return liked, nil
}

// TODO: just a stub
// POST https://api.twitter.com/1.1/statuses/destroy/:tweet-id.json
func (svc *TwitterImpl) DeleteTweets(user TwitterUser) (response interface{}, err error) {
	liked, err := svc.GetLikedTweets(user)
	if err != nil {
		log.Fatal(err)
	}
	log.Debug(liked)

	//TODO: implement undo like
	for k, v := range liked.Data {
		log.Debug(k, v)
	}

	// TODO: update to real response
	return liked, nil
}

// curl --location --request GET 'https://api.twitter.com/2/users/:user_id/tweets'
// https://documenter.getpostman.com/view/9956214/T1LMiT5U#daeb8a9f-6dac-4a40-add6-6b68bffb40cc
func (svc *TwitterImpl) GetReTweets(user TwitterUser) (tweets Tweets, err error) {
	if err != nil {
		log.Errorf("Error getting user:\n %v", err)
		return tweets, err
	}

	url := "https://api.twitter.com/2/users/" + user.Data.ID + "/tweets?referenced_tweets.type=retweeted"
	data, err := svc.request(url, http.MethodGet)
	if err != nil {
		log.Errorf("Error performing request:\n %v", err)
		return tweets, err
	}
	log.Debugf("GetReTweets: %s", data)

	if err = json.Unmarshal([]byte(*data), &tweets); err != nil {
		log.Error(err)
		return tweets, err
	}

	return tweets, nil
}

// POST https://api.twitter.com/1.1/statuses/unretweet/:tweet-id.json
func (svc *TwitterImpl) UndoReTweets(user TwitterUser) (response interface{}, err error) {
	tweets, err := svc.GetTweets(user)
	if err != nil {
		return nil, err
	}

	for k, v := range tweets.Data {
		if strings.Contains(v.Text, "RT @") {
			log.Debugf("delete retweet: %v", k)
		}
	}

	return nil, nil
}

// curl --location --request GET 'https://api.twitter.com/2/users/:user_id/tweets'
// https://documenter.getpostman.com/view/9956214/T1LMiT5U#daeb8a9f-6dac-4a40-add6-6b68bffb40cc
// TODO: stub
func (svc *TwitterImpl) GetReplies(user TwitterUser) (tweets Tweets, err error) {
	if err != nil {
		log.Errorf("Error getting user:\n %v", err)
		return tweets, err
	}

	url := "https://api.twitter.com/2/users/" + user.Data.ID + "/tweets?referenced_tweets.type=retweeted"
	data, err := svc.request(url, http.MethodGet)
	if err != nil {
		log.Errorf("Error performing request:\n %v", err)
		return tweets, err
	}
	log.Debugf("GetReTweets: %s", data)

	if err = json.Unmarshal([]byte(*data), &tweets); err != nil {
		log.Error(err)
		return tweets, err
	}

	return tweets, nil
}
