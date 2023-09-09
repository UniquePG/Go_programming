package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashingPassword(password string) (string, error) {

	// hash our password
	pass, err := bcrypt.GenerateFromPassword([]byte(password), 10)


	return string(pass), err	

}


func ComparePassword(password, hash string) bool {

	// compare the password it takes hashed password and original pass

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil

}

func main() {

	pass := "This is password";

	HashPass, _ := HashingPassword(pass)

	fmt.Println("Password: ", pass)
	fmt.Println("Hash Password: ", HashPass)

	result := ComparePassword(pass, HashPass)

	fmt.Println("Match: ", result)


}
