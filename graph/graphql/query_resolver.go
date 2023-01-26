package graphql

import (
	"context"

	"com.example/graphql/graph/model"
)

// Meetups is the resolver for the meetups field.
func (r *queryResolver) Meetups(ctx context.Context) ([]*model.Meetup, error) {
	return r.MeetupsRepo.GetMeetups()
	// return meetups, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	return r.UsersRepo.GetUserById(id)
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
