package main

import (
	"encoding/json"
	"fmt"
)

//type UserRole int16
//
//const (
//	UserRoleUnknown UserRole = 0
//	UserRoleAdmin   UserRole = 1
//	UserRoleModer   UserRole = 2
//	UserRoleUser    UserRole = 3
//)
//
//type User struct {
//	FirstName string
//	Age       int
//	Role      UserRole
//}

func main() {
	user := &User{
		FirstName: "Дима",
		Age:       20,
		Role:      UserRoleAdmin,
	}

	data, _ := json.MarshalIndent(user, "", "  ")
	fmt.Println(string(data))

	fmt.Println(user.StringRole())
}

func (u *User) StringRole() string {
	switch u.Role {
	case UserRoleAdmin:
		return "Admin"
	case UserRoleModer:
		return "Moder"
	case UserRoleUser:
		return "User"
	default:
		return "Unknown"
	}
}
