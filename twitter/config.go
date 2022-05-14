package twitter

/**
	It's the intention that this is where we deal with the Authentication and Authorization
**/

import (
	"flag"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gomodule/oauth1/oauth"
	log "github.com/sirupsen/logrus"
)

// authBearer only allows the app to read public information
var authBearer = flag.String("twitter.AuthBearer", "", "Twitter Authorization Bearer Token")
var username = flag.String("twitter.Username", "", "Twitter User Name")
var accessToken = flag.String("twitter.AccessToken", "", "Twitter Access Token")
var accessTokenSecret = flag.String("twitter.AccessTokenSecret", "", "Twitter Access Token Secret")
var apiKey = flag.String("twitter.ApiKey", "", "Twitter Consumer API Key")
var apiKeySecret = flag.String("twitter.ApiKeySecret", "", "Twitter Consumer API Secret")
var oAuthCallBackUrl = flag.String("twitter.OAuthCallBackUrl", "oob", "OAuth Call Back URL")
var exemptAuthors = flag.String("twitter.ExemptAuthors", "", "Exempt authors from (re)tweet, responses, and (un)likes")
var archive = flag.String("twitter.Archive", "", "The twitter users downloaded archive path (already extracted)")

type Config struct {
	UserName           string
	AuthBearer         string
	OAuthCallBackUrl   string
	APICredentials     oauth.Credentials
	UserCredentials    oauth.Credentials
	TwitterExemptUsers []string
	Archive            string
}

var oauthClient = oauth.Client{
	TemporaryCredentialRequestURI: "https://api.twitter.com/oauth/request_token",
	ResourceOwnerAuthorizationURI: "https://api.twitter.com/oauth/authorize",
	TokenRequestURI:               "https://api.twitter.com/oauth/access_token",
}

func NewConfig() Config {
	log.WithField("Twitter NewConfig Runtime", time.Now()).Info("**** New Twitter Config ****")
	flag.Parse()
	config := Config{
		UserName:         *username,
		AuthBearer:       *authBearer,
		OAuthCallBackUrl: *oAuthCallBackUrl,
		APICredentials: oauth.Credentials{
			Token:  *apiKey,
			Secret: *apiKeySecret,
		},
		UserCredentials: oauth.Credentials{
			Token:  *accessToken,
			Secret: *accessTokenSecret,
		},
		TwitterExemptUsers: strings.Split(*exemptAuthors, ","),
		Archive:            *archive,
	}
	oauthClient.Credentials = oauth.Credentials{
		Token:  *apiKey,
		Secret: *apiKeySecret,
	}
	log.WithField("NewConfig", config).Debug()
	return config
}

func (config Config) AuthorizationBearerToken() http.Header {
	headers := http.Header{}
	if len(config.AuthBearer) > 0 {
		headers.Add("Authorization", "Bearer "+config.AuthBearer)
	}
	log.WithField("Headers", headers).Debug("AuthorizationBearerToken")
	return headers
}

func (config Config) OAuthTokens(method string, resource *url.URL, form url.Values) http.Header {
	head := http.Header{}
	oauthClient.SetAuthorizationHeader(head, &config.UserCredentials, method, resource, form)
	log.WithField("Headers", head).Debug("OAuthTokens")
	return head
}
