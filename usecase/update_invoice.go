package usecase

import (
	"github.com/rwirdemann/restvoice/foundation"
	"github.com/rwirdemann/restvoice/domain"
	"time"
)

type UpdateInvoiceRepository interface {
	UpdateInvoice(invoice *domain.Invoice)
	RateByProjectIdAndActivityId(projectId int, activityId int) domain.Rate
	ActivityById(user string, id int) domain.Activity
	GetBookingsByInvoiceId(invoiceId int) []domain.Booking
}

type UpdateInvoice struct {
	invoiceConsumer   foundation.Consumer
	invoiceIdConsumer foundation.Consumer
	repository        UpdateInvoiceRepository
}

func NewUpdateInvoice(invoiceConsumer foundation.Consumer, invoiceIdConsumer foundation.Consumer, repository UpdateInvoiceRepository) *UpdateInvoice {
	return &UpdateInvoice{
		repository:        repository,
		invoiceConsumer:   invoiceConsumer,
		invoiceIdConsumer: invoiceIdConsumer}
}

func (u UpdateInvoice) Run(i ...interface{}) interface{} {
	invoice := u.invoiceConsumer.Consume(i[0]).(*domain.Invoice)
	invoice.Id = u.invoiceIdConsumer.Consume(i[1]).(int)

	time.Sleep(1*time.Second)
	// Aggregate positions
	if invoice.Status == "payment expected" {
		bookings := u.repository.GetBookingsByInvoiceId(invoice.Id)
		for _, b := range bookings {
			activity := u.repository.ActivityById("ralf", b.ActivityId)
			rate := u.repository.RateByProjectIdAndActivityId(b.ProjectId, b.ActivityId)
			invoice.AddPosition(b.ProjectId, activity.Name, b.Hours, rate.Price)
		}
	}
	u.repository.UpdateInvoice(invoice)
	return nil
}

func (UpdateInvoice) Cancel() {
	panic("implement me")
}

