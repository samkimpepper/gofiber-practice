package repository

import (
	"context"
	"go-note/ent"
	"go-note/ent/user"

	"github.com/google/uuid"
)

type UserRepository interface {
	FindByID(userID uuid.UUID) (*ent.User, error)
}

type userRepository struct {
	db *ent.Client
}

func NewUserRepository(db *ent.Client) UserRepository {
	return &userRepository{
		db: db,
	}
}

// ========================================================

func (repo userRepository) FindByID(userID uuid.UUID) (*ent.User, error) {

	selectedUser, err := repo.db.User.Query().
		Where(user.ID(userID)).
		Only(context.TODO())
	if err != nil {
		return nil, err
	}

	return selectedUser, nil
}
