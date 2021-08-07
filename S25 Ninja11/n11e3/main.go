package main

import "log"

type customErr struct {
	description string
}

func (ce customErr) Error() string {
	return ce.description
}

func foo(e error) {
	log.Println(e)
}

func main() {
	ce := customErr{
		description: "This is the description of the custom error",
	}

	foo(ce)
}
