package usecase

import (
	"github.com/rwirdemann/restvoice/foundation"
	"github.com/rwirdemann/restvoice/domain"
)

type GetInvoiceRepository interface {
	GetInvoice(id int, join ...string) domain.Invoice
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
	id := u.invoiceIdConsumer.Consume(i[0]).(int)
	join := u.joinConsumer.Consume(i[1]).(string)

	// Fetch invoice from database, join additional data
	invoice := u.repository.GetInvoice(id, join)

	return u.presenter.Present(invoice)
}

func (GetInvoice) Cancel() {
	panic("implement me")
}
