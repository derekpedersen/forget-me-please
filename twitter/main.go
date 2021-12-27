package twitter

import (
	"flag"
	"fmt"

	"github.com/derekpedersen/forget-me-please/utilities"
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
	twts, err := NewTweets(auth, user)
	if err != nil {
		log.Error(err)
	}
	fmt.Println("\nTwitter Options:")
	fmt.Println("Un(L)ike")
	fmt.Println("Un(R)eTweet")
	fmt.Println("Un(T)weet")
	char := utilities.ReadLine()
	switch *char {
	case "L":
		err = twts.Unlike()
		if err != nil {
			log.Error(err)
		}
	case "R":
		err = twts.UnRetweet()
		if err != nil {
			log.Error(err)
		}
	case "T":
		err = twts.Delete()
		if err != nil {
			log.Error(err)
		}
	default:
		fmt.Println("Please select an option")
	}
}
