package database

import "github.com/rwirdemann/restvoice/domain"

type MySQLRepository struct {
	nextId   int
	invoices map[int]*domain.Invoice
	bookings map[int]map[int]domain.Booking
}

func NewMySQLRepository() *MySQLRepository {
	r := MySQLRepository{}
	r.invoices = make(map[int]*domain.Invoice)
	r.bookings = make(map[int]map[int]domain.Booking)
	return &r
}

func (r *MySQLRepository) GetBookingsByInvoiceId(id int) []domain.Booking {
	var bookings []domain.Booking
	for _, b := range r.bookings[id] {
		bookings = append(bookings, b)
	}
	return bookings
}

func (r *MySQLRepository) GetInvoice(id int) domain.Invoice {
	return *r.invoices[id]
}

func (r *MySQLRepository) CreateInvoice(invoice *domain.Invoice) domain.Invoice {
	invoice.Id = r.getNextId()
	invoice.Status = "open"
	r.invoices[invoice.Id] = invoice
	return *invoice
}

func (r *MySQLRepository) CreateBooking(booking domain.Booking) {
	booking.Id = r.getNextId()
	if bookings, ok := r.bookings[booking.InvoiceId]; ok {
		bookings[booking.Id] = booking
	} else {
		bookings := make(map[int]domain.Booking)
		bookings[booking.Id] = booking
		r.bookings[booking.InvoiceId] = bookings
	}
}

func (r *MySQLRepository) getNextId() int {
	r.nextId = r.nextId + 1;
	return r.nextId
}
