package main

import (
	"fmt"
)

type AdminService interface {
	GetAdminInfo(string) (*Admin, error)
}

type impl struct{}

func (impl) GetAdminInfo(targetAdminName string) (*Admin, error) {
	fmt.Println(targetAdminName)
	return &Admin{
		Name: "gaku",
		Age:  "26",
	}, nil
}

type Admin struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}
