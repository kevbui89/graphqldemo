package graphql

import (
	"context"
	"errors"
	"fmt"

	"com.example/graphql/graph/model"
)

type mutationResolver struct{ *Resolver }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// CreateMeetup is the resolver for the createMeetup field.
func (r *mutationResolver) CreateMeetup(ctx context.Context, input model.NewMeetup) (*model.Meetup, error) {
	if len(input.Name) < 3 {
		return nil, errors.New("name not long enough")
	}

	if len(input.Description) < 3 {
		return nil, errors.New("description not long enough")
	}

	meetup := &model.Meetup{
		Name:        input.Name,
		Description: input.Description,
		UserID:      "1",
	}

	return r.MeetupsRepo.CreateMeetup(meetup)
}

// UpdateMeetup is the resolver for the updateMeetup field.
func (r *mutationResolver) UpdateMeetup(ctx context.Context, id string, input model.UpdateMeetup) (*model.Meetup, error) {
	meetup, err := r.MeetupsRepo.GetById(id)
	fmt.Println("meetup: ", meetup.ID)
	if err != nil || meetup == nil {
		return nil, errors.New("meetup does not exist")
	}

	didUpdate := false

	if input.Name != nil {
		if len(*input.Name) < 3 {
			return nil, errors.New("name is not long enough")
		}
		meetup.Name = *input.Name
		didUpdate = true
	}

	if input.Description != nil {
		if len(*input.Description) < 3 {
			return nil, errors.New("description is not long enough")
		}
		meetup.Description = *input.Description
		didUpdate = true
	}

	if !didUpdate {
		return nil, errors.New("no update done")
	}

	meetup, err = r.MeetupsRepo.Update(meetup)
	if err != nil {
		return nil, fmt.Errorf("error while updating meetup: %v", err)
	}

	return meetup, nil
}

// DeleteMeetup is the resolver for the deleteMeetup field.
func (r *mutationResolver) DeleteMeetup(ctx context.Context, id string) (bool, error) {
	meetup, err := r.MeetupsRepo.GetById(id)
	fmt.Println("meetup: ", meetup.ID)
	if err != nil || meetup == nil {
		return false, errors.New("meetup does not exist")
	}

	err = r.MeetupsRepo.Delete(meetup)
	if err != nil {
		return false, fmt.Errorf("error while deleting meetup: %v", err)
	}

	return true, nil
}
