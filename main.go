package main

import (
	"cat/console/repl"
	"fmt"
	"github.com/Grubba27/painter"
	"os"
	"os/user"
)

func main() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! Welcome to the Cat lang REPL 🐱\n", u.Username)
	t := paint.InPurple("Some text \n")
	fmt.Printf(t)
	repl.Start(os.Stdin, os.Stdout)
}
