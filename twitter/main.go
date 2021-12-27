package twitter

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	flag.Parse()
	log.SetLevel(log.DebugLevel)

	TakeAction()
}

func TakeAction() {
	auth := NewAuth()
	user, err := NewUser(auth)
	if err != nil {
		log.Error(err)
	}
	fmt.Println("Twitter Options:")
	fmt.Println("Un(L)ike")
	fmt.Println("Un(R)eTweet")
	fmt.Println("Un(T)weet")
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		log.Error(err)
	}
	switch char {
	case 'L':
		// twts, err := NewTweetsLiked(auth, user, "liked_tweets")
		// if err != nil {
		// 	log.Error(err)
		// }
		// err = twts.Unlike()
		// if err != nil {
		// 	log.Error(err)
		// }
	case 'R':
		twts, err := NewTweets(auth, user)
		if err != nil {
			log.Error(err)
		}
		err = twts.UnRetweet()
		if err != nil {
			log.Error(err)
		}
	case 'T':
		twts, err := NewTweets(auth, user)
		if err != nil {
			log.Error(err)
		}
		err = twts.Delete()
		if err != nil {
			log.Error(err)
		}
	default:
		fmt.Println("Please select an option")
	}
}
