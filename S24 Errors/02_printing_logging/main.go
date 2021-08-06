package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")

	if err != nil {
		fmt.Println("fmt.Println Err happened", err)
		log.Println("log.Println Err happened", err)

		fmt.Println("Panicln")
		log.Panicln(err)
		panic(err)

		// fmt.Println("Fatalln")
		//log.Fatalln(err)

	}
}
