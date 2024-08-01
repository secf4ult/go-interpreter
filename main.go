package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s, this is the Monkey Programming Language!\n", usr.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
