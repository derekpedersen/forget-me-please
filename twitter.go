package main

// https://developer.twitter.com/en/docs/authentication/oauth-1-0a/pin-based-oauth

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

// authBearer only allows the app to read public information
var twitterAuthBearer = flag.String("twitterAuthBearer", "", "Twitter Authorization Bearer Token")
var twitterUsername = flag.String("twitterUsername", "", "Twitter User Name")
var twitterAccessToken = flag.String("twitterAccessToken", "", "Twitter Access Token")
var twitterAccessTokenSecret = flag.String("twitterAccessTokenSecret", "", "Twitter Access Token Secret")
var twitterApiKey = flag.String("twitterApiKey", "", "Twitter API Key")
var twitterApiKeySecret = flag.String("twitterApiKeySecret", "", "Twitter API Secret")
var twitterOAuthCallBackUrl = flag.String("twitterOAuthCallBackUrl", "oob", "OAuth Call Back URL")

// {"data":{"id":"1684445455","name":"Derek Pedersen","username":"PedersenDerek"}}
type TwitterUser struct {
	Data struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		UserName string `json:"username"`
	} `json:"data"`
}

// {"data":[{"id":"1426928113426993152","text":"All of you guys are making fun of Marianne Williamson because you don't have any better ideas. Let's give it a shot. Let's deploy Jimmy Dore to Afghanistan."}]
type Tweets struct {
	Data []Tweet `json:"data"`
}

type Tweet struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

// Twitter interface
type Twitter interface {
	GetUser() (TwitterUser, error)
	GetReplies(user TwitterUser) (tweets Tweets, err error)
	GetTweets(user TwitterUser) (tweets Tweets, err error)
	GetReTweets(user TwitterUser) (tweets Tweets, err error)
	GetLikedTweets(user TwitterUser) (tweets Tweets, err error)
	UnLikeTweet(user TwitterUser, tweet Tweet) (response interface{}, err error)
	UnlikeTweets(user TwitterUser) (response interface{}, err error)
	DeleteTweets(user TwitterUser) (response interface{}, err error)
	UndoReTweet(user TwitterUser, tweet Tweet) (response interface{}, err error)
	UndoReTweets(user TwitterUser) (response interface{}, err error)
}

// TwitterImpl struct
type TwitterImpl struct {
}

// NewTwitter creates a new album service
func NewTwitter(authToken, accessToken, accessTokenSecret string) Twitter {
	return &TwitterImpl{}
}

func (svc *TwitterImpl) userTokens() {

}

// --header 'authorization: OAuth oauth_consumer_key="CONSUMER_API_KEY", oauth_nonce="OAUTH_NONCE", oauth_signature="OAUTH_SIGNATURE", oauth_signature_method="HMAC-SHA1", oauth_timestamp="OAUTH_TIMESTAMP", oauth_token="ACCESS_TOKEN", oauth_version="1.0"' \
func (svc *TwitterImpl) request(url, method string) (*string, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Errorf("Error creating request:\n %v", err)
		return nil, err
	}
	req.Header.Add("Accept", "application/json")

	if len(*twitterAuthBearer) > 0 {
		var bearer = "Bearer " + *twitterAuthBearer
		fmt.Print(bearer)
		req.Header.Add("Authorization", bearer)
	}

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

// GetUser is mainly used to the the ID of the user, something that is not displayed via the UI
func (svc *TwitterImpl) GetUser() (TwitterUser, error) {
	var twitterUser TwitterUser

	url := "https://api.twitter.com/2/users/by/username/" + *twitterUsername
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

// GetTweets gets the entire collection of tweets for a user (in paged format)
func (svc *TwitterImpl) GetTweets(user TwitterUser) (tweets Tweets, err error) {
	url := "https://api.twitter.com/2/users/" + user.Data.ID + "/tweets"
	data, err := svc.request(url, http.MethodGet)
	if err != nil {
		log.Errorf("Error performing request:\n %v", err)
		return tweets, err
	}
	log.Debugf("GetTweets: %v", data)

	if err = json.Unmarshal([]byte(*data), &tweets); err != nil {
		log.Error(err)
		return tweets, err
	}

	return tweets, nil
}

// GetTweets gets the entire collection of liked tweets for a user (in paged format)
func (svc *TwitterImpl) GetLikedTweets(user TwitterUser) (tweets Tweets, err error) {
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
func (svc *TwitterImpl) UnLikeTweet(user TwitterUser, tweet Tweet) (response interface{}, err error) {
	url := "https://api.twitter.com/2/users/" + user.Data.ID + "/likes/" + tweet.ID
	data, err := svc.request(url, http.MethodDelete)
	if err != nil {
		log.Errorf("Error performing request:\n %v", err)
		return nil, err
	}
	log.Debugf("Unlike Tweet: %v", data)

	if err = json.Unmarshal([]byte(*data), &response); err != nil {
		log.Error(err)
		return nil, err
	}

	return response, nil
}

func (svc *TwitterImpl) UnlikeTweets(user TwitterUser) (response interface{}, err error) {
	liked, err := svc.GetLikedTweets(user)
	if err != nil {
		log.Fatal(err)
	}
	log.Debug(liked)

	for _, v := range liked.Data {
		_, err := svc.UnLikeTweet(user, v)
		if err != nil {
			log.Fatal(err)
		}
	}

	return true, nil
}

// TODO: not working
// POST https://api.twitter.com/1.1/statuses/destroy/:tweet-id.json
func (svc *TwitterImpl) DeleteTweet(user TwitterUser, tweet Tweet) (response interface{}, err error) {
	url := "https://api.twitter.com/1.1/statuses/destroy/" + tweet.ID + ".json"
	data, err := svc.request(url, http.MethodPost)
	if err != nil {
		log.Errorf("Error performing request:\n %v", err)
		return nil, err
	}
	log.Debugf("Delete Tweet: %v", data)

	if err = json.Unmarshal([]byte(*data), &response); err != nil {
		log.Error(err)
		return nil, err
	}

	return response, nil
}

func (svc *TwitterImpl) DeleteTweets(user TwitterUser) (response interface{}, err error) {
	tweets, err := svc.GetTweets(user)
	if err != nil {
		log.Fatal(err)
	}
	log.Debug(tweets)

	for _, v := range tweets.Data {
		svc.DeleteTweet(user, v)
	}

	return true, nil
}

// TODO: not working
// POST https://api.twitter.com/1.1/statuses/unretweet/:tweet-id.json
func (svc *TwitterImpl) UndoReTweet(user TwitterUser, tweet Tweet) (response interface{}, err error) {
	url := "https://api.twitter.com/1.1/statuses/unretweet/" + tweet.ID + ".json"
	data, err := svc.request(url, http.MethodPost)
	if err != nil {
		log.Errorf("Error performing request:\n %v", err)
		return nil, err
	}
	log.Debugf("Undo Re Tweet: %v", data)

	if err = json.Unmarshal([]byte(*data), &response); err != nil {
		log.Error(err)
		return nil, err
	}

	return response, nil
}

func (svc *TwitterImpl) UndoReTweets(user TwitterUser) (response interface{}, err error) {
	tweets, err := svc.GetTweets(user)
	if err != nil {
		return nil, err
	}

	for _, v := range tweets.Data {
		if strings.Contains(v.Text, "RT @") {
			_, err := svc.UndoReTweet(user, v)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	return nil, nil
}

// TODO: not working
// curl --location --request GET 'https://api.twitter.com/2/users/:user_id/tweets'
// https://documenter.getpostman.com/view/9956214/T1LMiT5U#daeb8a9f-6dac-4a40-add6-6b68bffb40cc
func (svc *TwitterImpl) GetReplies(user TwitterUser) (tweets Tweets, err error) {
	url := "https://api.twitter.com/2/users/" + user.Data.ID + "/tweets" //?referenced_tweets.type=retweeted"
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

// TODO: not working
// curl --location --request GET 'https://api.twitter.com/2/users/:user_id/tweets'
// https://documenter.getpostman.com/view/9956214/T1LMiT5U#daeb8a9f-6dac-4a40-add6-6b68bffb40cc
func (svc *TwitterImpl) GetReTweets(user TwitterUser) (tweets Tweets, err error) {
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
