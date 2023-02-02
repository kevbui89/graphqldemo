package domain

import "com.example/graphql/graph/postgres"

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
