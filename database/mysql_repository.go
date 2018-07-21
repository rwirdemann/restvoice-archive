package database

import "github.com/rwirdemann/restvoice/domain"

type MySQLRepository struct {
	invoices map[int]*domain.Invoice
}

func NewMySQLRepository() *MySQLRepository {
	r := MySQLRepository{}
	r.invoices = make(map[int]*domain.Invoice)
	return &r
}

func (r *MySQLRepository) Invoices() []*domain.Invoice {
	var invoices []*domain.Invoice
	for _, invoice := range r.invoices {
		invoices = append(invoices, invoice)
	}
	return invoices
}

func (r *MySQLRepository) GetInvoice(id int) domain.Invoice {
	return *r.invoices[id]
}

func (r *MySQLRepository) CreateInvoice(invoice *domain.Invoice) {
	invoice.Id = r.nextId()
	invoice.Status = "open"
	r.invoices[invoice.Id] = invoice
}

func (r *MySQLRepository) nextId() int {
	nextId := 1
	for _, i := range r.invoices {
		if i.Id >= nextId {
			nextId = i.Id + 1
		}
	}
	return nextId
}
