package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	pwd := []byte("password")

	bs, err := bcrypt.GenerateFromPassword(pwd, 4)
	if err != nil {
		fmt.Println("err")
	}

	fmt.Println("Pwd:       ", string(pwd))
	fmt.Println("Hashed Pwd:", string(bs))

	err = bcrypt.CompareHashAndPassword(bs, []byte("password"))
	if err != nil {
		fmt.Println("Wrong password")
	} else {
		fmt.Println("Success!")
	}

}
