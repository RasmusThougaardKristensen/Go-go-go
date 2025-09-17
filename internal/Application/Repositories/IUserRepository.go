package IRepositories

import "GoSupplyChain/internal/Domain/models"

type IUserRepository interface {
	GetUserById(id string) (*models.UserModel, error)
	CreateUser(user *models.UserModel) error
}
