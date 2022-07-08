package twitter

import (
	"encoding/json"
	"net/http"

	"github.com/derekpedersen/forget-me-please/utilities"
	log "github.com/sirupsen/logrus"
)

// {"data":{"id":"1684445455","name":"Derek Pedersen","username":"PedersenDerek"}}
type User struct {
	Data struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		UserName string `json:"username"`
	} `json:"data"`
}

// GetUser is mainly used to the the ID of the user, something that is not displayed via the UI
func NewUser(config Config) (User, error) {
	var user User
	url := "https://api.twitter.com/2/users/by/username/" + config.UserName
	data, err := utilities.HttpRequest(url, http.MethodGet, config.AuthorizationBearerToken())
	if err != nil {
		log.Errorf("Error performing request:\n %v", err)
		return user, err
	}
	if err = json.Unmarshal([]byte(*data), &user); err != nil {
		log.Error(err)
		return user, err
	}
	return user, nil
}
