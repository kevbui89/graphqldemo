package postgres

import (
	"fmt"

	"com.example/graphql/graph/model"
	"github.com/go-pg/pg/v10"
)

type MeetupsRepo struct {
	DB *pg.DB
}

func (m *MeetupsRepo) GetMeetups(filter *model.MeetupFilter, limit, offset *int) ([]*model.Meetup, error) {
	var meetups []*model.Meetup

	query := m.DB.Model(&meetups).Order("id")

	if filter != nil {
		if filter.Name != nil && *filter.Name != "" {
			query.Where("name ILIKE ?", fmt.Sprintf("%%%s%%", *filter.Name))
		}
	}

	if limit != nil {
		query.Limit(*limit)
	}

	if offset != nil {
		query.Offset(*offset)
	}

	err := query.Select()
	if err != nil {
		return nil, err
	}
	return meetups, nil
}

func (m *MeetupsRepo) CreateMeetup(meetup *model.Meetup) (*model.Meetup, error) {
	_, err := m.DB.Model(meetup).Returning("*").Insert()
	if err != nil {
		return &model.Meetup{}, nil
	}
	return meetup, err
}

func (m *MeetupsRepo) GetById(id string) (*model.Meetup, error) {
	var meetup model.Meetup
	err := m.DB.Model(&meetup).Where("id = ?", id).First()
	return &meetup, err
}

func (m *MeetupsRepo) Update(meetup *model.Meetup) (*model.Meetup, error) {
	_, err := m.DB.Model(meetup).Where("id = ?", meetup.ID).Update()
	return meetup, err
}

func (m *MeetupsRepo) Delete(meetup *model.Meetup) error {
	_, err := m.DB.Model(meetup).Where("id = ?", meetup.ID).Delete()
	return err
}

func (m *MeetupsRepo) GetMeetupsForUser(user *model.User) ([]*model.Meetup, error) {
	var meetups []*model.Meetup
	err := m.DB.Model(&meetups).Where("user_id = ?", user.ID).Order("id").Select()
	return meetups, err
}
