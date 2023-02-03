package domain

import (
	"errors"

	"com.example/graphql/graph/model"
	"com.example/graphql/graph/postgres"
)

var (
	ErrBadCredentials  = errors.New("email/password combination does not work")
	ErrUnauthenticated = errors.New("unauthenticated")
	ErrForbidden       = errors.New("unauthorized")
)

type Domain struct {
	UsersRepo   postgres.UserRepo
	MeetupsRepo postgres.MeetupsRepo
}

func NewDomain(usersRepo postgres.UserRepo, meetupsRepo postgres.MeetupsRepo) *Domain {
	return &Domain{
		UsersRepo:   usersRepo,
		MeetupsRepo: meetupsRepo,
	}
}

type Ownable interface {
	IsOwner(user *model.User) bool
}

func checkOwnership(o Ownable, user *model.User) bool {
	return o.IsOwner(user)
}
