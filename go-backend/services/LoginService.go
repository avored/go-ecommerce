package services

import (
	"fmt"
)



func LoginAdminUser (email string, password string) {
	fmt.Printf("Login %s USER %s", email, password)
}
