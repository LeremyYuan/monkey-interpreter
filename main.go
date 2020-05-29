package main

import (
	"log"
	"monkey-interpreter/repl"
	"os"
	user2 "os/user"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	user, err := user2.Current()
	if err != nil {
		panic(err)
	}
	log.Printf("Hello %s! This is the Monkey programming language!\n",
		user.Username)
	log.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)

}
