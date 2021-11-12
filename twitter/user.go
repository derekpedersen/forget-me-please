package twitter

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type User struct {
	Data struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		UserName string `json:"username"`
	} `json:"data"`
}

// GetUser is mainly used to the the ID of the user, something that is not displayed via the UI
func NewUser(auth Auth) (User, error) {
	var user User

	url := "https://api.twitter.com/2/users/by/username/" + auth.UserName
	data, err := httpRequest(url, http.MethodGet, auth.AuthorizationBearerToken())
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
