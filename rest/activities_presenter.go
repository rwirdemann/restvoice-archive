package rest

import (
	"encoding/json"
	"time"
	"github.com/rwirdemann/restvoice/domain"
)

type CacheableActivities struct {
	Activities   []byte
	LastModified time.Time
}

type ActivitiesPresenter struct {
}

func NewActivitiesPresenter() ActivitiesPresenter {
	return ActivitiesPresenter{}
}

func (j ActivitiesPresenter) Present(i interface{}) interface{} {
	lastModified := time.Unix(0, 0)
	activities := i.([]domain.Activity)
	for _, a := range activities {
		if a.Updated.After(lastModified) {
			lastModified = a.Updated
		}
	}
	b, _ := json.Marshal(i)
	return CacheableActivities{Activities: b, LastModified: lastModified}
}
