package main

import (
	"fmt"
	"os"
	"os/user"
	"sanaa/repl"
)

func main() {
	user, err := user.Current()

	execPath, _ := os.Executable()
	hostName, _ := os.Hostname()

	fmt.Println("EXEC PATH:", execPath)
	fmt.Println("HOSTNAME :", hostName)
	fmt.Println("")

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Sanaa programming language!\n",
		user.Username)
	fmt.Printf("Feel free to type in comands\n")

	repl.Start(os.Stdin, os.Stdout)
}
