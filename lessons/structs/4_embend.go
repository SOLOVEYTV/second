package main

import (
	"encoding/json"
	"fmt"
)

type (
	UserRole int16
	City     int32
)

const (
	UserRoleUnknown UserRole = 0
	UserRoleAdmin   UserRole = 1
	UserRoleModer   UserRole = 2
	UserRoleUser    UserRole = 3

	CityUnknown City = 0
	CityMoscow  City = 1
	CitySpb     City = 2
)

type HomeAddress struct {
	City        City
	Street      string
	HouseNumber int
}

type User struct {
	FirstName string
	Age       int
	Role      UserRole
	*HomeAddress
}

func main() {
	user := &User{
		FirstName: "Дима",
		Age:       20,
		Role:      UserRoleAdmin,
		HomeAddress: &HomeAddress{
			City:        CityMoscow,
			Street:      "Садовая",
			HouseNumber: 10,
		},
	}

	data, _ := json.MarshalIndent(user, "", "    ")
	fmt.Println(string(data))

	fmt.Println(user.StringCity()) // вызываем метод встроенной структуры
}

func (a *HomeAddress) StringCity() string {
	switch a.City {
	case CityMoscow:
		return "Moscow"
	case CitySpb:
		return "Spb"
	default:
		return "Unknown"
	}
}
