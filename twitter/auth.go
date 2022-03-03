package twitter

import (
	"flag"
	"net/http"
	"net/url"
	"strings"

	"github.com/gomodule/oauth1/oauth"
)

// authBearer only allows the app to read public information
var twitterAuthBearer = flag.String("twitterAuthBearer", "", "Twitter Authorization Bearer Token")
var twitterUsername = flag.String("twitterUsername", "", "Twitter User Name")
var twitterAccessToken = flag.String("twitterAccessToken", "", "Twitter Access Token")
var twitterAccessTokenSecret = flag.String("twitterAccessTokenSecret", "", "Twitter Access Token Secret")
var twitterApiKey = flag.String("twitterApiKey", "", "Twitter Consumer API Key")
var twitterApiKeySecret = flag.String("twitterApiKeySecret", "", "Twitter Consumer API Secret")
var twitterOAuthCallBackUrl = flag.String("twitterOAuthCallBackUrl", "oob", "OAuth Call Back URL")
var twitterExemptUsers = flag.String("twitterExemptUsers", "", "Twitter users whose (Re)Tweets you want to keep")

type TwitterAuth struct {
	UserName           string
	AuthBearer         string
	AccessToken        string
	AccessTokenSecret  string
	ApiKey             string
	ApiKeySecret       string
	OAuthCallBackUrl   string
	TwitterExemptUsers []string
}

var oauthClient = oauth.Client{
	TemporaryCredentialRequestURI: "https://api.twitter.com/oauth/request_token",
	ResourceOwnerAuthorizationURI: "https://api.twitter.com/oauth/authorize",
	TokenRequestURI:               "https://api.twitter.com/oauth/access_token",
}

func NewTwitterAuth() TwitterAuth {
	oauthClient.Credentials = oauth.Credentials{
		Token:  *twitterApiKey,
		Secret: *twitterApiKeySecret,
	}
	return TwitterAuth{
		UserName:           *twitterUsername,
		AuthBearer:         *twitterAuthBearer,
		AccessToken:        *twitterAccessToken,
		AccessTokenSecret:  *twitterAccessTokenSecret,
		ApiKey:             *twitterApiKey,
		ApiKeySecret:       *twitterApiKeySecret,
		OAuthCallBackUrl:   *twitterOAuthCallBackUrl,
		TwitterExemptUsers: strings.Split(*twitterExemptUsers, ","),
	}
}

func (auth TwitterAuth) AuthorizationBearerToken() http.Header {
	headers := http.Header{}
	if len(auth.AuthBearer) > 0 {
		headers.Add("Authorization", "Bearer "+auth.AuthBearer)
	}
	return headers
}

func (auth TwitterAuth) OAuthTokens(method string, resource *url.URL, form url.Values) http.Header {
	head := http.Header{}

	userCredentials := oauth.Credentials{
		Token:  *twitterAccessToken,
		Secret: *twitterAccessTokenSecret,
	}

	oauthClient.SetAuthorizationHeader(head, &userCredentials, method, resource, form)

	return head
}
