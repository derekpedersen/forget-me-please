package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// {"data":{"id":"1684445455","name":"Derek Pedersen","username":"PedersenDerek"}}
type TwitterUser struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"username"`
}

// Twitter interface
type Twitter interface {
	GetLikedTweets() (json string, err error)
}

// TwitterImpl struct
type TwitterImpl struct {
	authToken string
	userName  string
}

// NewTwitter creates a new album service
func NewTwitter(authToken, userName string) *TwitterImpl {
	return &TwitterImpl{
		authToken: authToken,
		userName:  userName,
	}
}

func (svc *TwitterImpl) GetUser() (user TwitterUser, err error) {
	url := "https://api.twitter.com/2/users/by/username/" + svc.userName

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Errorf("Error creating request:\n %v", err)
		return user, err
	}
	req.Header.Add("Accept", "application/json")

	// Create a Bearer string by appending string access token
	var bearer = "Bearer " + svc.authToken

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Errorf("Error making request:\n %v", err)
		return user, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Errorf("Error reading res.Body:\n %v", err)
		return user, err
	}

	data := string(body)

	if err = json.Unmarshal([]byte(data), &user); err != nil {
		log.Error(err)
		return user, err
	}

	return user, nil
}

// curl --location --request GET 'https://api.twitter.com/2/users//liked_tweets'
func (svc *TwitterImpl) GetLikedTweets(userID string) (json string, err error) {

	url := "https://api.twitter.com/2/users/" + userID + "/liked_tweets"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Errorf("Error creating request:\n %v", err)
		return "", err
	}
	req.Header.Add("Accept", "application/json")

	// Create a Bearer string by appending string access token
	var bearer = "Bearer " + svc.authToken

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Errorf("Error making request:\n %v", err)
		return "", err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Errorf("Error reading res.Body:\n %v", err)
		return "", err
	}

	json = string(body)

	log.Debugf("GetLikedTweets: %s", json)

	return json, nil
}
