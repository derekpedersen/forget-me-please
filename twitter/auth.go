package twitter

import (
	"flag"
	"net/http"
)

// authBearer only allows the app to read public information
var twitterAuthBearer = flag.String("twitterAuthBearer", "", "Twitter Authorization Bearer Token")
var twitterUsername = flag.String("twitterUsername", "", "Twitter User Name")
var twitterAccessToken = flag.String("twitterAccessToken", "", "Twitter Access Token")
var twitterAccessTokenSecret = flag.String("twitterAccessTokenSecret", "", "Twitter Access Token Secret")
var twitterApiKey = flag.String("twitterApiKey", "", "Twitter API Key")
var twitterApiKeySecret = flag.String("twitterApiKeySecret", "", "Twitter API Secret")
var twitterOAuthCallBackUrl = flag.String("twitterOAuthCallBackUrl", "oob", "OAuth Call Back URL")

type TwitterAuth struct {
	UserName          string
	AuthBearer        string
	AccessToken       string
	AccessTokenSecret string
	ApiKey            string
	ApiKeySecret      string
	OAuthCallBackUrl  string
}

func NewTwitterAuth() TwitterAuth {
	return TwitterAuth{
		UserName:          *twitterUsername,
		AuthBearer:        *twitterAuthBearer,
		AccessToken:       *twitterAccessToken,
		AccessTokenSecret: *twitterAccessTokenSecret,
		ApiKey:            *twitterApiKey,
		ApiKeySecret:      *twitterApiKeySecret,
		OAuthCallBackUrl:  *twitterOAuthCallBackUrl,
	}
}

func (auth TwitterAuth) SetAuthorizationBearerToken() http.Header {
	headers := http.Header{}
	if len(auth.AuthBearer) > 0 {
		headers.Add("Authorization", "Bearer "+auth.AuthBearer)
	}
	return headers
}

func (auth TwitterAuth) SetOAuthTokens() http.Header {
	headers := http.Header{}
	if len(*twitterAuthBearer) > 0 {
		var bearer = "Bearer " + auth.AuthBearer
		headers.Add("Authorization", bearer)
	}
	return headers
}
