package usecase

import (
	"github.com/rwirdemann/restvoice/foundation"
	"github.com/rwirdemann/restvoice/domain"
	"strings"
)

type GetInvoiceRepository interface {
	GetInvoice(id int) domain.Invoice
	GetBookingsByInvoiceId(id int) []domain.Booking
}

type GetInvoice struct {
	invoiceIdConsumer foundation.Consumer
	joinConsumer      foundation.Consumer
	presenter         foundation.Presenter
	repository        GetInvoiceRepository
}

func NewGetInvoice(consumer foundation.Consumer, expandConsumer foundation.Consumer,
	presenter foundation.Presenter, repository GetInvoiceRepository) *GetInvoice {
	return &GetInvoice{
		repository:        repository,
		invoiceIdConsumer: consumer,
		joinConsumer:      expandConsumer,
		presenter:         presenter,
	}
}

func (u GetInvoice) Run(i ...interface{}) interface{} {
	// Fetch invoice from database
	id := u.invoiceIdConsumer.Consume(i[0]).(int)
	invoice := u.repository.GetInvoice(id)

	// Join additional data
	join := u.joinConsumer.Consume(i[0]).(string)
	if strings.Contains(join, "bookings") {
		invoice.Bookings = u.repository.GetBookingsByInvoiceId(id)
	}

	return u.presenter.Present(invoice)
}
