package domain

import (
	"fmt"
)

type Position struct {
	Hours float32
	Price float32
}

type Invoice struct {
	Id         int                         `json:"id"`
	Status     string                      `json:"status"`
	CustomerId int                         `json:"customerId"`
	Year       int                         `json:"year"`
	Month      int                         `json:"month"`
	Positions  map[int]map[string]Position `json:"positions"`
}

func (i *Invoice) AddPosition(projectId int, title string, hours float32, rate float32) {
	if i.Positions == nil {
		i.Positions = make(map[int]map[string]Position)
	}

	if i.Positions[projectId] == nil {
		i.Positions[projectId] = make(map[string]Position)
	}

	if p, ok := i.Positions[projectId][title]; ok {
		p.Hours = p.Hours + hours
		p.Price = p.Price + hours*rate
		i.Positions[projectId][title] = p
	} else {
		i.Positions[projectId][title] = Position{Hours: hours, Price: hours * rate}
	}
}

func (i *Invoice) String() string {
	s := ""
	for k, p := range i.Positions {
		s = fmt.Sprintf("%s\nProjekt: %d", s, k)
		for activity, hoursAndPrice := range p {
			s = fmt.Sprintf("%s\n %s: %.2f => %2.f", s, activity, hoursAndPrice.Hours, hoursAndPrice.Price)
		}
	}
	return s
}
