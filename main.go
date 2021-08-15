package main

import (
	"flag"
	"fmt"
	"log"
)

var bearer = flag.String("bearer", "", "Authorization Bearer Token")
var username = flag.String("username", "", "Twitter User Name")

func main() {
	twitter := NewTwitter(*bearer, *username)

	user, err := twitter.GetUser()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)

	liked, err := twitter.GetLikedTweets(user.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(liked)
}
