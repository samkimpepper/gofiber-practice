package repository

import (
	"context"
	"go-note/ent"
	"go-note/ent/user"
	"go-note/module/auth/entities"
	"log"

	"github.com/pkg/errors"
)

type AuthRepository interface {
	Save(dto entities.RegisterRequest) (*ent.User, error)
}

type authRepository struct {
	db *ent.Client
}

func NewAuthRepository(db *ent.Client) AuthRepository {
	return authRepository{db: db}
}

func (repo authRepository) Save(dto entities.RegisterRequest) (*ent.User, error) {

	isExists, err := repo.db.User.Query().
		Where(user.Email(dto.Email)).
		Exist(context.TODO()) // 이거 find 말고 exists 없나?
	if isExists {
		log.Println("authRepository error: %v\n", err)
		return nil, entities.ErrUserAlreadyExists
	}

	saveduser, err := repo.db.User.Create().
		SetEmail(dto.Email).
		SetUsername(dto.Username).
		SetPassword(dto.Password).
		Save(context.TODO())
	if err != nil {
		return nil, errors.Wrap(err, "repository.authRepository.Save()")
	}

	return saveduser, nil
}

func (repo authRepository) FindByEmail(email string) (*ent.User, error) {
	selectedUser, err := repo.db.User.Query().
		Where(user.Email(email)).
		Only(context.TODO())
	if err != nil {
		return nil, err
	}

	return selectedUser, nil
}
