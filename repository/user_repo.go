package repository

import (
	"context"
	"timewise/model"
)

type UserRepo interface {

	SaveUser(context context.Context , user model.User) (model.User, error)
	SelectUserByEmail(conetxt context.Context, email string ) (model.User , error)
}