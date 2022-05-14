package twitter

/**
	It's the intention that this is where we deal with the parsing of archival data
**/

import (
	"encoding/json"
	"io/ioutil"
)

type likedTweet struct {
	tweetId     string
	fullText    string
	expandedUrl string
}

type like struct {
	like likedTweet
}

type liked []like

func newArchivedTweets(config Config, likedTweets *bool) (tweets Tweets, err error) {
	// TODO: make this actually read a file
	return tweets, nil
}

func newLiked(filepath string) (liked liked, err error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return liked, err
	}
	err = json.Unmarshal(data, &liked)
	if err != nil {
		return liked, err
	}

	return liked, nil
}

func (liked *liked) parseTweets() (tweets []Tweet, err error) {
	for _, v := range *liked {
		t := Tweet{
			ID: v.like.tweetId,
		}
		tweets = append(tweets, t)
	}
	return tweets, nil
}
