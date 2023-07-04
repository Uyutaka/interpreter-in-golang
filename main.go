package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/uyutaka/interpreter-in-golang/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! Thisis the Monkey programming langeuage!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
