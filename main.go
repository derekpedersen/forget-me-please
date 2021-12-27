package main

import (
	"fmt"

	"github.com/derekpedersen/forget-me-please/twitter"
)

func main() {
	fmt.Println("(F)acebook")
	fmt.Println("(R)eddit")
	fmt.Println("(T)witter")
	fmt.Println("(Y)ouTube")
	fmt.Println("(I)nstagram")
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	// text := scanner.Text()
	// switch text {
	// case "T":
	twitter.TakeAction()
	// default:
	// 	fmt.Println("Not yet supported sorry")
	// }
}
