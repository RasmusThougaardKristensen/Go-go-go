package Repositories

import (
	"GoSupplyChain/internal/Application/Repositories"
	"GoSupplyChain/internal/Domain/models"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type UserRepository struct {
	database *sqlx.DB
}

func PostgresUserRepository(database *sqlx.DB) IRepositories.IUserRepository {
	return &UserRepository{database: database}
}

func (userRepository *UserRepository) GetUserById(id string) (*models.UserModel, error) {
	var entity UserEntity

	if _, err := uuid.Parse(id); err != nil {
		return nil, fmt.Errorf("invalid UUID format: %s", id)
	}

	query := `SELECT id, name, email FROM users WHERE id = $1`

	err := userRepository.database.Get(&entity, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with id %s not found", id)
		}
		return nil, fmt.Errorf("failed to get user by id %s: %w", id, err)
	}

	return mapToModel(entity), nil
}

func mapToEntity(user models.UserModel) *UserEntity {
	return &UserEntity{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}
}

func mapToModel(user UserEntity) *models.UserModel {
	return &models.UserModel{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}
}
