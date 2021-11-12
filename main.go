package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/derekpedersen/forget-me-please/twitter"
)

func main() {
	fmt.Println("Please select a platform from below")
	fmt.Println("-----------------------------------")
	fmt.Println("F: Facebook")
	fmt.Println("R: Reddit")
	fmt.Println("T: Twitter")
	fmt.Println("Y: YouTube")
	fmt.Println("I: Instagram")
	fmt.Println("-----------------------------------")
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
	}
	switch char {
	case 'T':
		twitter.TakeAction()
	default:
		fmt.Println("Not yet supported sorry")
	}
}
