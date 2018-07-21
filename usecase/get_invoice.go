package usecase

import (
	"github.com/rwirdemann/restvoice/foundation"
	"github.com/rwirdemann/restvoice/domain"
)

type GetInvoiceRepository interface {
	GetInvoice(id int) domain.Invoice
}

type GetInvoice struct {
	invoiceIdConsumer foundation.Consumer
	presenter         foundation.Presenter
	repository        GetInvoiceRepository
}

func NewGetInvoice(consumer foundation.Consumer, presenter foundation.Presenter, repository GetInvoiceRepository) *GetInvoice {
	return &GetInvoice{
		repository:        repository,
		invoiceIdConsumer: consumer,
		presenter:         presenter,
	}
}

func (u GetInvoice) Run(i ...interface{}) interface{} {
	id := u.invoiceIdConsumer.Consume(i[0]).(int)
	return u.presenter.Present(u.repository.GetInvoice(id))
}
