package twitter

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type Tweet struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

func (twt *Tweet) Unlike(auth TwitterAuth, user TwitterUser) error {
	url := "https://api.twitter.com/2/users/" + user.Data.ID + "/likes/" + twt.ID
	data, err := httpRequest(url, http.MethodDelete, auth.SetOAuthTokens())
	if err != nil {
		log.Errorf("Error performing request:\n %v", err)
		return err
	}
	log.Debugf("Unlike Tweet: %v", data)

	var response interface{}
	if err = json.Unmarshal([]byte(*data), &response); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (twt *Tweet) Delete(auth TwitterAuth, user TwitterUser) error {
	url := "https://api.twitter.com/1.1/statuses/destroy/" + twt.ID + ".json"
	data, err := httpRequest(url, http.MethodPost, auth.SetOAuthTokens())
	if err != nil {
		log.Errorf("Error performing request:\n %v", err)
		return err
	}
	log.Debugf("Delete Tweet: %v", data)

	var response interface{}
	if err = json.Unmarshal([]byte(*data), &response); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (twt *Tweet) UnRetweet(auth TwitterAuth, user TwitterUser) error {
	url := "https://api.twitter.com/1.1/statuses/destroy/" + twt.ID + ".json"
	data, err := httpRequest(url, http.MethodPost, auth.SetOAuthTokens())
	if err != nil {
		log.Errorf("Error performing request:\n %v", err)
		return err
	}
	log.Debugf("Delete Tweet: %v", data)

	var response interface{}
	if err = json.Unmarshal([]byte(*data), &response); err != nil {
		log.Error(err)
		return err
	}

	return nil
}
