package services

import (
	"ALP/src/models"
	"fmt"
)

func CreateUser (user models.User) error {
	fmt.Printf("Creating user: %v\n", user)
	
	return nil


}