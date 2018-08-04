package usecase

import (
	"github.com/rwirdemann/restvoice/foundation"
	"github.com/rwirdemann/restvoice/domain"
)

type CreateBookingRepository interface {
	CreateBooking(booking domain.Booking) domain.Booking
}

type CreateBooking struct {
	invoiceIdConsumer foundation.Consumer
	bookingConsumer   foundation.Consumer
	presenter         foundation.Presenter
	repository        CreateBookingRepository
}

func NewCreateBooking(bookingConsumer foundation.Consumer, invoiceIdConsumer foundation.Consumer, presenter foundation.Presenter, repository CreateBookingRepository) *CreateBooking {
	return &CreateBooking{
		repository:        repository,
		bookingConsumer:   bookingConsumer,
		invoiceIdConsumer: invoiceIdConsumer,
		presenter:         presenter}
}

func (u CreateBooking) Run(i ...interface{}) interface{} {
	booking := u.bookingConsumer.Consume(i[0]).(*domain.Booking)
	booking.InvoiceId = u.invoiceIdConsumer.Consume(i[1]).(int)
	created := u.repository.CreateBooking(*booking)
	return u.presenter.Present(created)
}
