package main

import (
	"fmt"

	"github.com/alejandrodgb/learn-golang/S29-Ninja13/n13e2/quote"
	"github.com/alejandrodgb/learn-golang/S29-Ninja13/n13e2/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
