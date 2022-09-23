package repository

import (
	"context"
	"go-note/ent"
	"go-note/ent/user"
	"go-note/module/auth/entities"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

type AuthRepository interface {
	Save(dto entities.RegisterRequest) (*ent.User, error)
	FindByEmail(email string) (*ent.User, error)
}

type authRepository struct {
	db  *ent.Client
	rdb *redis.Client
}

func NewAuthRepository(db *ent.Client, rdb *redis.Client) AuthRepository {
	return authRepository{db: db, rdb: rdb}
}

func (repo authRepository) Save(dto entities.RegisterRequest) (*ent.User, error) {
	if repo.ExistsByEmail(dto.Email) {
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

func (repo authRepository) ExistsByEmail(email string) bool {
	isExists, _ := repo.db.User.Query().
		Where(user.Email(email)).
		Exist(context.TODO())

	return isExists
}

// func (repo authRepository) Logout(rt string) error {

// }
