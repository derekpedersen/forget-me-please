package twitter

/**
	It's the intention that this is where we deal with timeline based information
**/

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/derekpedersen/forget-me-please/utilities"
	log "github.com/sirupsen/logrus"
)

func newTimeline(config Config, user User, paginationToken *string, likedTweets *bool) (Tweets, error) {
	log.WithField("NewTweets", time.Now())
	var tweets Tweets
	url := "https://api.twitter.com/2/users/" + user.Data.ID
	if likedTweets != nil && *likedTweets {
		url += "/liked_tweets"
	} else {
		url += "/tweets"
	}
	url += "?max_results=100"
	// this line tries to avoid cached responses from twitter
	// url += "&" + utilities.Random() + "=" + utilities.Random()
	if paginationToken != nil && len(*paginationToken) > 0 {
		url += "&pagination_token=" + *paginationToken
	}
	data, err := utilities.HttpRequest(url, http.MethodGet, config.AuthorizationBearerToken())
	if err != nil {
		log.Errorf("Error performing request:\n %v", err)
		return tweets, err
	}
	log.WithFields(log.Fields{"Tweets": data}).Debug("NewTweets")

	if err = json.Unmarshal([]byte(*data), &tweets); err != nil {
		log.Error(err)
		return tweets, err
	}
	tweets.Config = config
	tweets.User = user
	log.WithFields(log.Fields{"Tweets": tweets, "API Response": data, "URL": url}).Debug("NewTweets")
	return tweets, nil
}
