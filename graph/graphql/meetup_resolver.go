package graphql

import (
	"context"

	"com.example/graphql/graph/model"
)

type meetupResolver struct{ *Resolver }

// UserID is the resolver for the user_id field.
func (r *meetupResolver) UserID(ctx context.Context, obj *model.Meetup) (*model.User, error) {
	return getUserLoader(ctx).Load(obj.UserID)
}

// Meetup returns MeetupResolver implementation.
func (r *Resolver) Meetup() MeetupResolver { return &meetupResolver{r} }
