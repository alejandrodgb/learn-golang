package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Open("/Users/adgb/go/src/github.com/alejandrodgb/learn-golang/S24 Errors/04_open_file/names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
