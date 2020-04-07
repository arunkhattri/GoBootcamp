package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		*files = append(*files, path)
		return nil
	}
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Give me directory name to traverse.")
		return
	}
	var files []string

	err := filepath.Walk(args[0], visit(&files))
	if err != nil {
		fmt.Println(err)
		return
	}

	var items []byte

	for _, file := range files {
		items = append(items, file...)
		items = append(items, '\n')
	}

	error := ioutil.WriteFile("dirTree.txt", items, 0644)
	if error != nil {
		fmt.Println(err)
		return
	}

}
