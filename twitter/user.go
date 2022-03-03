package twitter

import (
	"encoding/json"
	"net/http"

	"github.com/derekpedersen/forget-me-please/utilities"
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

// GetUser is mainly used to the the ID of the user, something that is not displayed via the UI
func NewTwitterUser(auth TwitterAuth) (TwitterUser, error) {
	var twitterUser TwitterUser

	url := "https://api.twitter.com/2/users/by/username/" + auth.UserName
	data, err := utilities.HttpRequest(url, http.MethodGet, auth.AuthorizationBearerToken())
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
