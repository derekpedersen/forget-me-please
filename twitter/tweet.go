package twitter

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/derekpedersen/forget-me-please/utilities"
	log "github.com/sirupsen/logrus"
)

type Tweet struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

func (twt *Tweet) IsRetweet() bool {
	log.WithField("Tweet IsRetweet Runtime", time.Now())
	return strings.Contains(twt.Text, "RT @")
}

func (twt *Tweet) IsExempt(exempt []string) bool {
	log.WithField("Tweet IsExempt Runtime", time.Now())
	for _, v := range exempt {
		if strings.Contains(twt.Text, v) {
			return true
		}
	}
	return false
}

func (twt *Tweet) Unlike(auth Auth, user User) error {
	log.WithField("Tweet Unlike Runtime", time.Now())
	resource, _ := url.Parse("https://api.twitter.com/2/users/" + user.Data.ID + "/likes/" + twt.ID)
	// resource, _ := url.Parse("https://api.twitter.com/1.1/favorites/destroy.json?id=" + twt.ID)
	data, err := utilities.HttpRequest(resource.String(), http.MethodDelete, auth.OAuthTokens(http.MethodDelete, resource, nil))
	if err != nil {
		log.Errorf("Error performing request:\n %v", err)
		return err
	}
	log.WithFields(log.Fields{"ID": twt.ID, "Text": twt.Text, "API Response": data}).Printf("Unliked")

	var response interface{}
	if err = json.Unmarshal([]byte(*data), &response); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (twt *Tweet) Delete(auth Auth, user User) error {
	log.WithField("Tweet Delete Runtime", time.Now())
	resource, _ := url.Parse("https://api.twitter.com/1.1/statuses/destroy/" + twt.ID + ".json")
	data, err := utilities.HttpRequest(resource.String(), http.MethodPost, auth.OAuthTokens(http.MethodPost, resource, nil))
	if err != nil {
		log.Errorf("Error performing request:\n %v", err)
		return err
	}
	log.WithFields(log.Fields{"ID": twt.ID, "Text": twt.Text, "API Response": data}).Printf("Delete")

	var response interface{}
	if err = json.Unmarshal([]byte(*data), &response); err != nil {
		log.Error(err)
		return err
	}

	return nil
}
