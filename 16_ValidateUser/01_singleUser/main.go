package main

// Validate single user
// if username and password is incorrect display message
// Access denied for <username>
// if invalid password, display "Invalid password for <username>."
// if valid username and password, display
// Access granted for "<username>".

import (
	"fmt"
	"os"
	"strings"
)

const (
	msg      = `Usage: [username] [password]`
	errUser  = "Access denied for %q.\n"
	errPwd   = "Invalid password for %q.\n"
	accessOK = "Access granted to %q.\n"

	usrname = "arun"
	passwd  = "5771"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Println(strings.TrimSpace(msg))
		return
	}

	username, password := args[1], args[2]

	if username != usrname {
		fmt.Printf(errUser, username)
	} else if password != passwd {
		fmt.Printf(errPwd, username)
	} else {
		fmt.Printf(accessOK, username)
	}
}
