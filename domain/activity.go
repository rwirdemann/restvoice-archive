package domain

import "time"

type Activity struct {
	Id      int
	Name    string    `json:"name"`
	UserId  string    `json:"userId"` // belongs to user
	Updated time.Time `json:"-"`
}
