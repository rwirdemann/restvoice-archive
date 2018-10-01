package usecase

import (
	"github.com/rwirdemann/restvoice/foundation"
	"github.com/rwirdemann/restvoice/domain"
)

type CreateInvoiceRepository interface {
	CreateInvoice(invoice *domain.Invoice) domain.Invoice
}

type CreateInvoice struct {
	invoiceConsumer foundation.Consumer
	presenter       foundation.Presenter
	repository      CreateInvoiceRepository
}

func NewCreateInvoice(invoiceConsumer foundation.Consumer, presenter foundation.Presenter, repository CreateInvoiceRepository) *CreateInvoice {
	return &CreateInvoice{
		repository:      repository,
		invoiceConsumer: invoiceConsumer,
		presenter:       presenter}
}

func (u CreateInvoice) Run(i ...interface{}) interface{} {
	invoice := u.invoiceConsumer.Consume(i[0]).(*domain.Invoice)
	created := u.repository.CreateInvoice(invoice)
	return u.presenter.Present(created)
}

func (CreateInvoice) Cancel() {
	panic("implement me")
}
