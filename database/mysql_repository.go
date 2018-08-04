package database

import (
	"github.com/rwirdemann/restvoice/domain"
	"strings"
	"time"
)
import _ "github.com/go-sql-driver/mysql"

type MySQLRepository struct {
	nextId     int
	invoices   map[int]*domain.Invoice
	bookings   map[int]map[int]domain.Booking
	projects   map[int]domain.Project
	customers  map[int]domain.Customer
	activities map[string]map[int]domain.Activity
}

func (r *MySQLRepository) GetActivities(userId string) []domain.Activity {
	var activities []domain.Activity
	for _, a := range r.activities[userId] {
		activities = append(activities, a)
	}
	return activities
}

func NewMySQLRepository() *MySQLRepository {
	r := MySQLRepository{}
	r.invoices = make(map[int]*domain.Invoice)
	r.bookings = make(map[int]map[int]domain.Booking)
	r.projects = make(map[int]domain.Project)
	r.customers = make(map[int]domain.Customer)
	r.activities = make(map[string]map[int]domain.Activity)
	return &r
}

func (r *MySQLRepository) GetBookingsByInvoiceId(id int) []domain.Booking {
	var bookings []domain.Booking
	for _, b := range r.bookings[id] {
		bookings = append(bookings, b)
	}
	return bookings
}

func (r *MySQLRepository) GetInvoice(id int, join ...string) domain.Invoice {
	i := *r.invoices[id]
	if len(join) > 0 {
		if strings.Contains(join[0], "bookings") {
			i.Bookings = r.GetBookingsByInvoiceId(id)
		}
	}
	return i
}

func (r *MySQLRepository) GetProject(id int) domain.Project {
	return r.projects[id]
}

func (r *MySQLRepository) GetCustomer(id int) domain.Customer {
	return r.customers[id]
}

func (r *MySQLRepository) CreateInvoice(invoice *domain.Invoice) domain.Invoice {
	invoice.Id = r.getNextId()
	invoice.Status = "open"
	r.invoices[invoice.Id] = invoice
	return *invoice
}

func (r *MySQLRepository) CreateBooking(booking domain.Booking) domain.Booking {
	booking.Id = r.getNextId()
	if bookings, ok := r.bookings[booking.InvoiceId]; ok {
		bookings[booking.Id] = booking
	} else {
		bookings := make(map[int]domain.Booking)
		bookings[booking.Id] = booking
		r.bookings[booking.InvoiceId] = bookings
	}
	return booking
}

func (r *MySQLRepository) CreateActivity(activity domain.Activity) {
	activity.Id = r.getNextId()
	activity.Updated = time.Now().UTC()
	if activities, ok := r.activities[activity.UserId]; ok {
		activities[activity.Id] = activity
	} else {
		activities := make(map[int]domain.Activity)
		activities[activity.Id] = activity
		r.activities[activity.UserId] = activities
	}
}

func (r *MySQLRepository) getNextId() int {
	r.nextId = r.nextId + 1
	return r.nextId
}
