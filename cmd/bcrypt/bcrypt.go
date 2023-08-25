package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	// for i, arg := range os.Args {
	// 	fmt.Println(i, arg)
	// }

	switch os.Args[1] {
	case "hash":
		hash(os.Args[2])
	case "compare":
		compare(os.Args[2], os.Args[3])
	default:
		fmt.Printf("Invalid command: %v\n", os.Args[1])
	}
}

func hash(passorwd string) {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(passorwd), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("Error hashing %v\n", passorwd)
		return
	}

	fmt.Println(string(hashBytes))
}

func compare(password, hash string) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		fmt.Printf("password does not match: %v\n", password)
		return
	}

	fmt.Println("You're connected!")
}
