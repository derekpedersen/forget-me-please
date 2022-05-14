package twitter

/**
	It's the intention that this is where we deal with the parsing of archival data
**/

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

func newArchive(config Config, likedTweets *bool) (tweets Tweets, err error) {
	// TODO: make this actually read a file
	return tweets, nil
}

type likedTweet struct {
	tweetId     string
	fullText    string
	expandedUrl string
}

type like struct {
	like likedTweet
}

type liked []like

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

type archivedTweet struct {
	Retweeted bool
	Source    string
	// entities
	// display_text_range
	FavoriteCount string `json:"favorite_count"`
	IdStr         string `json:"id_str"`
	Truncated     bool
	RetweetCount  string    `json:"retweet_count"`
	ID            string    `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	Favorited     bool
	FullText      string `json:"full_text"`
	Lang          string
}

func newArchivedTweets(filepath string) (tweets []archivedTweet, err error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return tweets, err
	}
	err = json.Unmarshal(data, &tweets)
	if err != nil {
		return tweets, err
	}

	return tweets, nil
}
