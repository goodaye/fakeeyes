package service

import "fmt"

var (
	ErrorUserNotFound = fmt.Errorf("user not found")
	ErrorUserExist    = fmt.Errorf("user has been existed")
)
