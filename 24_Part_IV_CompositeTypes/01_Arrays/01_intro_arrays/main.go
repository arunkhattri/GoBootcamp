package main

import "fmt"

const (
	winter = 1
	summer = 3
	yearly = winter + summer
)

func main() {
	var books [yearly]string

	books[0] = "Anna Kareina"
	books[1] = "Madame Bovary"
	books[2] = "War and Peace"
	// books[3] = "The Great Gatsby"
	books[3] += books[0] + " 2nd Edition"

	// // print type
	// fmt.Printf("books : %T\n", books)
	// // print array
	// fmt.Printf("books : %q\n", books)
	// print array type and array
	fmt.Printf("books : %#v\n", books)

	var (
		wBooks [winter]string
		sBooks [summer]string
	)

	wBooks[0] = books[0]

	// for i := 0; i < len(sBooks); i++ {
	//	sBooks[i] = books[i+1]
	// }

	for i := range sBooks {
		sBooks[i] = books[i+1]
	}

	fmt.Printf("wBooks : %#v\n", wBooks)
	fmt.Printf("sBooks : %#v\n", sBooks)

	var published [len(books)]bool

	published[0] = true
	published[len(books)-1] = true

	fmt.Println("\nPublished Books:")
	for i, ok := range published {
		if ok {
			fmt.Printf("+ %s\n", books[i])
		}
	}

}
