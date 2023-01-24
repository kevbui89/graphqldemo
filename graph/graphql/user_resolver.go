package graphql

import (
	"context"

	"com.example/graphql/graph/model"
)

type userResolver struct{ *Resolver }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

// Meetups is the resolver for the meetups field.
func (r *userResolver) Meetups(ctx context.Context, obj *model.User) ([]*model.Meetup, error) {
	// var m []*model.Meetup
	// for _, mu := range meetups {
	// 	if mu.ID == obj.ID {
	// 		m = append(m, mu)
	// 	}
	// }

	return nil, nil
}
