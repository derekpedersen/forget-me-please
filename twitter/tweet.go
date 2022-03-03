package twitter

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/derekpedersen/forget-me-please/utilities"
	log "github.com/sirupsen/logrus"
)

type Tweet struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

func (twt *Tweet) IsRetweet() bool {
	return strings.Contains(twt.Text, "RT @")
}

func (twt *Tweet) IsExempt(exempt []string) bool {
	for _, v := range exempt {
		if strings.Contains(twt.Text, v) {
			return true
		}
	}
	return false
}

func (twt *Tweet) Unlike(auth Auth, user User) error {
	resource, _ := url.Parse("https://api.twitter.com/2/users/" + user.Data.ID + "/likes/" + twt.ID)
	data, err := utilities.HttpRequest(resource.String(), http.MethodDelete, auth.OAuthTokens(http.MethodDelete, resource, nil))
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

func (twt *Tweet) Delete(auth Auth, user User) error {
	resource, _ := url.Parse("https://api.twitter.com/1.1/statuses/destroy/" + twt.ID + ".json")
	data, err := utilities.HttpRequest(resource.String(), http.MethodPost, auth.OAuthTokens(http.MethodPost, resource, nil))
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
