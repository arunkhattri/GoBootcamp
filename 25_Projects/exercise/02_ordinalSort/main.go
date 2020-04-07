package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Send me some item and I'll sort them.")
		return
	}

	// sort the slice of strings
	sort.Strings(args)

	// write to a file
	var items []byte

	for i, item := range args {
		items = strconv.AppendInt(items, int64(i+1), 10)
		items = append(items, ". "...)
		items = append(items, item...)
		items = append(items, '\n')
	}

	err := ioutil.WriteFile("items.txt", items, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

}
