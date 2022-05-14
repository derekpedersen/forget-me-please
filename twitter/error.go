package twitter

import (
	"encoding/json"
	"net/url"

	"github.com/derekpedersen/forget-me-please/utilities"
	log "github.com/sirupsen/logrus"
)

type Response struct {
	Title  string
	Detail string
	Type   string
	Status int
	Errors []ErrorInfo
}

func (res *Response) NoStatusFound() bool {
	if res.Status == 144 {
		return true
	}
	if res.Errors != nil && len(res.Errors) > 0 && res.Errors[0].Code == 144 {
		return true
	}
	return false
}

type ErrorInfo struct {
	Code    int
	Message string
}

func Update(resource *url.URL, methodType string) (response Response, err error) {
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
