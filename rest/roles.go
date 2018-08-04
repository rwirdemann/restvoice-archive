package rest

import (
	"net/http"
	"github.com/rwirdemann/restvoice/domain"
)

type RoleRepository interface {
	GetProject(id int) domain.Project
	GetCustomer(id int) domain.Customer
}

func AssertOwnsCustomer(next http.HandlerFunc, repository RoleRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := extractJwtFromHeader(r.Header)
		jsonConsumer := NewJSONConsumer(domain.Booking{})
		booking := jsonConsumer.Consume(r.Body).(domain.Booking)
		if ownsCustomer(token, booking.ProjectId, repository) {
			next.ServeHTTP(w, r)
			return
		}
		w.WriteHeader(http.StatusForbidden)
	}
}

func ownsCustomer(token string, projectId int, repository RoleRepository) bool {
	userId := claim(token, "sub")
	project := repository.GetProject(projectId)
	customer := repository.GetCustomer(project.CustomerId)
	return customer.UserId == userId
}

