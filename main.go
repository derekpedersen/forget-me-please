package main

import (
	"fmt"

	"github.com/derekpedersen/forget-me-please/twitter"
	"github.com/derekpedersen/forget-me-please/utilities"
)

func main() {
	// TODO: need an option to go with a CLI or self hosted PWA

	fmt.Println("(F)acebook")
	fmt.Println("(R)eddit")
	fmt.Println("(T)witter")
	fmt.Println("(Y)ouTube")
	fmt.Println("(I)nstagram")

	text := utilities.ReadLine()
	switch *text {
	case "T":
		twitter.TakeAction()
	default:
		fmt.Println("Not yet supported sorry")
	}
}
