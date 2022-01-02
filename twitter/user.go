package twitter

import (
	"encoding/json"
	"net/http"

	"github.com/derekpedersen/forget-me-please/utilities"
	log "github.com/sirupsen/logrus"
)

type User struct {
	Data struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		UserName string `json:"username"`
	} `json:"data"`
}

func NewUser(auth Auth) (User, error) {
	var user User
	url := "https://api.twitter.com/2/users/by/username/" + auth.UserName
	data, err := utilities.HttpRequest(url, http.MethodGet, auth.AuthorizationBearerToken())
	if err != nil {
		log.Errorf("Error performing request:\n %v", err)
		return user, err
	}
	if err = json.Unmarshal([]byte(*data), &user); err != nil {
		log.Error(err)
		return user, err
	}
	log.WithField("User", user).Debug("NewUser")
	return user, nil
}
