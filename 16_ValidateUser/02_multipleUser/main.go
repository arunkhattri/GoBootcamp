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

	usr, usr2       = "arun", "jack"
	passwd, passwd2 = "5771", "1234"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Println(strings.TrimSpace(msg))
		return
	}

	username, password := args[1], args[2]

	if username != usr && username != usr2 {
		fmt.Printf(errUser, username)
	} else if username == usr && password == passwd {
		fmt.Printf(accessOK, username)
	} else if username == usr2 && password == passwd2 {
		fmt.Printf(accessOK, username)
	} else {
		fmt.Printf(errPwd, username)
	}
}
