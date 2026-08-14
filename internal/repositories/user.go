package repositories

import (
	"database/sql"
	"go-landing-page/internal/database"
	"go-landing-page/internal/database/queries"
	"go-landing-page/internal/dto"
)

type UserRepository struct {
	db database.Service
}

func NewUserRepository(db database.Service) UserRepository {
	return UserRepository{db: db}
}

func (r *UserRepository) Create(createUserReq dto.CreateUserReq) (sql.Result, error) {
	user, err := r.db.GetDB().Exec(queries.CreateUser, createUserReq)
	if err != nil {
		return nil, err
	}

	return user, nil
}
