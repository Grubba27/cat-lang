package main

import (
	"cat/console/color"
	"cat/console/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! Welcome to the Cat lang REPL 🐱\n", u.Username)
	fmt.Printf(
		color.Colorize("Start typing some code to see awesome things \n",
			color.Purple))
	repl.Start(os.Stdin, os.Stdout)
}
