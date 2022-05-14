package twitter

import (
	"encoding/json"
	"io/ioutil"
)

type archivedLikedTweet struct {
	tweetId     string
	fullText    string
	expandedUrl string
}

type archivedLike struct {
	like archivedLikedTweet
}

type archivedLiked struct {
	liked []archivedLike
}

func newArchiveLiked(filepath string) (liked archivedLiked, err error) {
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

func (liked *archivedLiked) parseTweets() (tweets []Tweet, err error) {
	for _, v := range liked.liked {
		t := Tweet{
			ID: v.like.tweetId,
		}
		tweets = append(tweets, t)
	}
	return tweets, nil
}
