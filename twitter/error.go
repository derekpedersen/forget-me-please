package twitter

import (
	"encoding/json"
	"net/url"

	"github.com/derekpedersen/forget-me-please/utilities"
	log "github.com/sirupsen/logrus"
)

type errorResponse struct {
	Title  string
	Detail string
	Type   string
	Status int
	Errors []twitterError
}

func (twt *errorResponse) NoStatusFound() bool {
	if twt.Status == 144 {
		return true
	}
	if twt.Errors != nil && len(twt.Errors) > 0 && twt.Errors[0].Code == 144 {
		return true
	}
	return false
}

type twitterError struct {
	Code    int
	Message string
}

func performRequest(resource *url.URL, methodType string) (response errorResponse, err error) {
	data, err := utilities.HttpRequest(resource.String(), methodType, config.OAuthTokens(methodType, resource, nil))
	if err != nil {
		log.Errorf("Error performing request:\n %v", err)
		return response, err
	}
	if err = json.Unmarshal([]byte(*data), &response); err != nil {
		log.Error(err)
		return response, err
	}
	log.WithFields(log.Fields{"Resource": resource, "HTTP Method": methodType, "API Response": *data, "Action Response": response}).Printf("Action")
	return response, nil
}
