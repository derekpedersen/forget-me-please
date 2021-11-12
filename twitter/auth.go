package twitter

import (
	"flag"
	"net/http"

	"github.com/gomodule/oauth1/oauth"
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

var oauthClient = oauth.Client{
	TemporaryCredentialRequestURI: "https://api.twitter.com/oauth/request_token",
	ResourceOwnerAuthorizationURI: "https://api.twitter.com/oauth/authorize",
	TokenRequestURI:               "https://api.twitter.com/oauth/access_token",
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

func (auth TwitterAuth) AuthorizationBearerToken() http.Header {
	headers := http.Header{}
	if len(auth.AuthBearer) > 0 {
		headers.Add("Authorization", "Bearer "+auth.AuthBearer)
	}
	return headers
}

func (auth TwitterAuth) OAuthTokens() http.Header {
	client := oauthClient
	client.Credentials = oauth.Credentials{
		Token:  *twitterAccessToken,
		Secret: *twitterApiKeySecret,
	}
	oauthClient.SetAuthorizationHeader()

	return oauthClient.Header.Clone()
}

// 	resp, err := oauthClient.Get(nil, tokenCred,
// 		"https://api.twitter.com/1.1/statuses/home_timeline.json", nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer resp.Body.Close()
// 	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
// 		log.Fatal(err)
// 	}
// }
