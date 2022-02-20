package users

import (
	"go.uber.org/zap"
)

type (
	UserAPI struct {
		i   Interactor
		log *zap.Logger
	}
)

func NewUserAPI(i Interactor) *UserAPI {
	return &UserAPI{
		i:   i,
		log: zap.L().Named("users"),
	}
}
