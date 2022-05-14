package twitter

import (
	"fmt"
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

// https://api.twitter.com/1.1/favorites/create.json?id=TWEET_ID_TO_FAVORITE
func (twt *Tweet) Like(auth Auth, user User) error {
	resource, _ := url.Parse("https://api.twitter.com/1.1/favorites/create.json?id=" + twt.ID)
	response, err := performRequest(resource, http.MethodPost)
	if err != nil {
		log.Errorf("Error performing request:\n %v", err)
		return err
	}
	log.WithFields(log.Fields{"ID": twt.ID, "Text": twt.Text, "API Response": response}).Printf("Liked")
	return nil
}

func (twt *Tweet) Unlike(auth Auth, user User) error {
	// resource, _ := url.Parse("https://api.twitter.com/2/users/" + user.Data.ID + "/likes/" + twt.ID)
	// response, err := performRequest(resource, http.MethodDelete)
	resource, _ := url.Parse("https://api.twitter.com/1.1/favorites/destroy.json?id=" + twt.ID)
	response, err := performRequest(resource, http.MethodPost)
	if err != nil {
		log.Errorf("Error performing request:\n %v", err)
		return err
	}
	if response.NoStatusFound() {
		// No Status Found with that ID
		// This is a twitter api error that has yet to be resolved due to
		// their internal caching mechanisms.
		// To get around this we will need "like" and then "unlike" the tweet
		// this will result in people getting notifications of your interacting
		// with their tweets. So people find this odd, I obviously do not.
		utilities.Delay()
		err = twt.Like(auth, user)
		if err != nil {
			log.Errorf("Error performing request:\n %v", err)
			return err
		}

		// Try again to unlike the tweet but only once
		// maybe have a multiple retry in the future
		utilities.Delay()
		response, err = performRequest(resource, http.MethodPost)
		if err != nil {
			log.Errorf("Error performing request:\n %v", err)
			return err
		}
		if response.NoStatusFound() {
			log.WithField("Error", response.Errors).Println("Twitter Unlike Failure")
			return fmt.Errorf("Twitter API Error: %v", response)
		} else if response.Status > 400 {
			log.Error(response)
			return fmt.Errorf("Twitter API Error: %v", response)
		}

	} else if response.Status > 400 {
		log.Error(response)
		return fmt.Errorf("Twitter API Error: %v", response)
	}
	log.WithFields(log.Fields{"ID": twt.ID, "Text": twt.Text, "API Response": response}).Printf("Unliked")
	return nil
}

func (twt *Tweet) Delete(auth Auth, user User) error {
	log.WithField("Tweet Delete Runtime", time.Now())
	resource, _ := url.Parse("https://api.twitter.com/1.1/statuses/destroy/" + twt.ID + ".json")
	response, err := performRequest(resource, http.MethodPost)
	if err != nil {
		return err
	}
	if response.Status > 400 {
		log.Error(response)
		return fmt.Errorf("Twitter API Error: %v", response)
	}
	log.WithFields(log.Fields{"ID": twt.ID, "Text": twt.Text, "API Response": response}).Printf("Deleted")
	return nil
}
