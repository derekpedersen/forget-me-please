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

func (auth TwitterAuth) AuthorizationBearerToken() http.Header {
	headers := http.Header{}
	if len(auth.AuthBearer) > 0 {
		headers.Add("Authorization", "Bearer "+auth.AuthBearer)
	}
	return headers
}

func (auth TwitterAuth) GetOAuthTokens() http.Header {
	headers := http.Header{}
	if len(*twitterAuthBearer) > 0 {
		var bearer = "Bearer " + auth.AuthBearer
		headers.Add("Authorization", bearer)
	}
	return headers
}

// import "github.com/gomodule/oauth1/oauth"
//
// type Credentials struct {
// 	Token  string // Also known as consumer key or access token.
// 	Secret string // Also known as consumer secret or access token secret.
// }
//
// tempCred, err := oauthClient.RequestTemporaryCredentials(nil, "oob", nil)
// 	if err != nil {
// 		log.Fatal("RequestTemporaryCredentials:", err)
// 	}

// 	u := oauthClient.AuthorizationURL(tempCred, nil)

// 	fmt.Printf("1. Go to %s\n2. Authorize the application\n3. Enter verification code:\n", u)

// 	var code string
// 	fmt.Scanln(&code)

// 	tokenCred, _, err := oauthClient.RequestToken(nil, tempCred, code)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

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
