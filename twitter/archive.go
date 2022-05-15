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
	if *likedTweets {
		liked, err := newLiked(config.Archive)
		if err != nil {
			return tweets, err
		}
		return liked.parseTweets()
	}
	twts, err := newArchivedTweets(config.Archive)
	if err != nil {
		return
	}
	return twts.parseTweets()
}

type archived_like struct {
	like struct {
		tweetId     string
		fullText    string
		expandedUrl string
	}
}

type archived_liked []archived_like

func newLiked(filepath string) (liked archived_liked, err error) {
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

func (liked *archived_liked) parseTweets() (tweets Tweets, err error) {
	for _, v := range *liked {
		t := Tweet{
			ID: v.like.tweetId,
		}
		tweets.Data = append(tweets.Data, t)
	}
	return tweets, nil
}

type archived_tweet struct {
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

type archived_tweets []archived_tweet

func newArchivedTweets(filepath string) (tweets archived_tweets, err error) {
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

func (arc *archived_tweets) parseTweets() (tweets Tweets, err error) {
	for _, v := range *arc {
		t := Tweet{
			ID: v.ID,
		}
		tweets.Data = append(tweets.Data, t)
	}
	return tweets, nil
}
