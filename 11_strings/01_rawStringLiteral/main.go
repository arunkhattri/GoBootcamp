package main

import "fmt"

func main() {
	var s string
	s = `
<html>
    <body>"Hello World !"</body>
</html>`
	fmt.Println(s)
	fmt.Println("C:\\Dropbox\\my_file") // string literal
	fmt.Println(`C:\Dropbox\myfile`)    // raw string literal
}
