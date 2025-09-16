package models

type UserModel struct {
	Id    string
	Name  string
	Email string
}

var UserModels = []UserModel{
	{Id: "1", Name: "Rasmus", Email: "Rasmus@gmail.com"},
	{Id: "2", Name: "Christian", Email: "Christian@gmail.com"},
	{Id: "3", Name: "Laura", Email: "Laura@gmail.com"},
}
