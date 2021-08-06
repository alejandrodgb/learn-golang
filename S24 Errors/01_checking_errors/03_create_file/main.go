package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("/Users/adgb/go/src/github.com/alejandrodgb/learn-golang/S24 Errors/04_open_file/names.txt")
	if err != nil {
		fmt.Println("error: ", err)
	}
	defer f.Close()

	r := strings.NewReader("This is text written by the program.")

	io.Copy(f, r)
}
