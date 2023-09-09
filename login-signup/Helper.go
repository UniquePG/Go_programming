package main

import "fmt"

func findUserByEmail(email string) (User, error) {

	var user User

	err := Dbvar.Where("email = ?", email).First(&user).Error

	if err != nil {
		fmt.Printf("user not found for this email", err.Error)

		return user, err
	}

	return user, nil
}