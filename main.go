package main

import (
	"fmt"
	"log"
)

func main() {
	passLength := 12
	num, err := generatePassword(passLength)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(num)
}
