package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hey %s, Wellcome to the Interpreter!\n", user.Name)

	fmt.Println("Start typing your commands")

	repl.Start(os.Stdin, os.Stdout)
}
