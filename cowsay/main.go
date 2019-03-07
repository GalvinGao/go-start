package main

import "fmt"
import "strings"
import "os"

func main() {
	messageTmpl := ` %s
< %s >
 %s
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||`

	message := strings.Join(os.Args[1:], " ")
	if len(message) == 0 {
		message = "I love Golang!"
	}
	messageTop := strings.Repeat("_", len(message))
	messageBottom := strings.Repeat("-", len(message))
	_, err := fmt.Printf(messageTmpl, messageTop, message, messageBottom)
	if err != nil {
		fmt.Println("wee woo wee woo! Print String failed! (how is this gonna happen...")
	}
}
