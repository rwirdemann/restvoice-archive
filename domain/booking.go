package domain

type Booking struct {
	Id          int     `json:"id"`
	Day         int     `json:"day"`
	Hours       float32 `json:"hours"`
	Description string  `json:"description"`
	InvoiceId   int     `json:"invoiceId,omitempty"`  // belongs to invoice
	ProjectId   int     `json:"projectId,omitempty"`  // belongs to project
	ActivityId  int     `json:"activityId,omitempty"` // belongs to activity
}
