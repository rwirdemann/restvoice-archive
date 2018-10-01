package usecase

import (
	"github.com/rwirdemann/restvoice/foundation"
	"github.com/rwirdemann/restvoice/domain"
)

type GetActivitiesRepository interface {
	GetActivities(userId string) []domain.Activity
}

type GetActivities struct {
	userConsumer foundation.Consumer
	presenter    foundation.Presenter
	repository   GetActivitiesRepository
}

func NewGetActivities(consumer foundation.Consumer, presenter foundation.Presenter, repository GetActivitiesRepository) *GetActivities {
	return &GetActivities{
		userConsumer: consumer,
		repository:   repository,
		presenter:    presenter,
	}
}

func (u GetActivities) Run(i ...interface{}) interface{} {
	userId := u.userConsumer.Consume(i[0]).(string)
	activities := u.repository.GetActivities(userId)
	return u.presenter.Present(activities)
}

func (GetActivities) Cancel() {
	panic("implement me")
}

