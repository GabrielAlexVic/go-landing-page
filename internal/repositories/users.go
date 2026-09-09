package repositories

import (
	"database/sql"
	"go-landing-page/internal/database"
	"go-landing-page/internal/database/queries"
	"go-landing-page/internal/dto"
	"go-landing-page/internal/models"
)

type UserRepository struct {
	db database.Service
}

func NewUserRepository(db database.Service) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(createUserReq dto.CreateUserReq) (sql.Result, error) {
	user, err := r.db.GetDB().NamedExec(queries.CreateUser, createUserReq)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) GetByEmail(email string) (models.User, error) {
	var user models.User
	err := r.db.GetDB().Get(&user, queries.GetUserByEmail, email)
	if err != nil {
		return models.User{}, err
	}

	return user, err
}
