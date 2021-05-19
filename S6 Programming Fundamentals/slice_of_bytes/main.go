package main

import "fmt"

func main() {
	s := "Alejandro"

	fmt.Println(s)
	fmt.Printf("%T\n", s)

	// Convert a string to a slice of bytes
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("Slice of byes: %T\n\n", bs)

	// Print the UTF-8 codepoint
	// Reference https://pkg.go.dev/fmt
	fmt.Println("Printing the UTF-8 codepoint")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U", s[i])
	}

	// Print the Unicode code point from the string as decimal
	fmt.Println("\n\nPrint the Unicode code point from the string as decimal")
	for i, v := range s {
		fmt.Println(i, v)
	}

	// Print the Unicode code point from the string as decimal
	fmt.Println("\n\nPrint the Unicode code point from the string as Hex")
	for i, v := range s {
		fmt.Printf("at index position %d we have hex %#x\n", i, v)
	}
}
