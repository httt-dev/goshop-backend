package repository

import (
	"context"
	"timewise/model"
)

type UserRepo interface {
	SaveUser(context context.Context, user model.User) (model.User, error)
	SelectUserByEmail(conetxt context.Context, email string) (model.User, error)
	SelectUserById(conetxt context.Context, userId string) (model.User, error)
}
