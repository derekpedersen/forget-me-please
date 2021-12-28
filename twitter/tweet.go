package twitter

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/derekpedersen/forget-me-please/utilities"
	log "github.com/sirupsen/logrus"
)

type Tweet struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

func (twt *Tweet) Unlike(auth Auth, user User) error {
	resource, _ := url.Parse("https://api.twitter.com/2/users/" + user.Data.ID + "/likes/" + twt.ID)
	log.WithField("resource", resource)
	data, err := utilities.HttpRequest(resource.String(), http.MethodDelete, auth.OAuthTokens(http.MethodDelete, resource, nil))
	if err != nil {
		log.Errorf("Error performing request:\n %v", err)
		return err
	}
	log.WithField("data", data)
	var response interface{}
	if err = json.Unmarshal([]byte(*data), &response); err != nil {
		log.Error(err)
		return err
	}
	log.WithField("response", response).Debug("Unlike")
	return nil
}

func (twt *Tweet) Delete(auth Auth, user User) error {
	resource, _ := url.Parse("https://api.twitter.com/1.1/statuses/destroy/" + twt.ID + ".json")
	log.WithField("resource", resource)
	data, err := utilities.HttpRequest(resource.String(), http.MethodPost, auth.OAuthTokens(http.MethodPost, resource, nil))
	if err != nil {
		log.Errorf("Error performing request:\n %v", err)
		return err
	}
	log.WithField("data", data)
	var response interface{}
	if err = json.Unmarshal([]byte(*data), &response); err != nil {
		log.Error(err)
		return err
	}
	log.WithField("response", response).Debug("Delete")
	return nil
}

func (twt *Tweet) UnRetweet(auth Auth, user User) error {
	resource, _ := url.Parse("https://api.twitter.com/1.1/statuses/unretweet/" + twt.ID + ".json")
	log.WithField("resource", resource)
	data, err := utilities.HttpRequest(resource.String(), http.MethodPost, auth.OAuthTokens(http.MethodPost, resource, nil))
	if err != nil {
		log.Errorf("Error performing request:\n %v", err)
		return err
	}
	log.WithField("data", data)
	var response interface{}
	if err = json.Unmarshal([]byte(*data), &response); err != nil {
		log.Error(err)
		return err
	}
	log.WithField("response", response).Debug("UnRetweet")
	return nil
}
