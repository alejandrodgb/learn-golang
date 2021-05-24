package main

import (
	"fmt"
	"strconv"
)

func main() {
	hex := "0xffffff801edbf84"

	value, err := strconv.ParseInt(hex, 0, 0)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Hex = %s | Decimal = %d | Binary = %b\n", hex, value, value)
	}

}
