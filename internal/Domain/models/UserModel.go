package models

import "github.com/google/uuid"

type UserModel struct {
	Id    string
	Name  string
	Email string
}

func User(name, email string) UserModel {
	return UserModel{
		Id:    uuid.New().String(),
		Name:  name,
		Email: email,
	}
}

var UserModels = []UserModel{
	{Id: "1", Name: "Rasmus", Email: "Rasmus@gmail.com"},
	{Id: "2", Name: "Christian", Email: "Christian@gmail.com"},
	{Id: "3", Name: "Laura", Email: "Laura@gmail.com"},
}
