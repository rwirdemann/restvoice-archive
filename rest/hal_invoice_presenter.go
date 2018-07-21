package rest

import (
	"encoding/json"
	"fmt"
	"github.com/rwirdemann/restvoice/domain"
	"errors"
	"log"
)

type HALInvoicePresenter struct {
}

func NewHALInvoicePresenter() HALInvoicePresenter {
	return HALInvoicePresenter{}
}

type Link struct {
	Href string `json:"href"`
}

type HALDecorator struct {
	domain.Invoice
	Links    map[string]Link `json:"_links"`
	Embedded interface{}     `json:"_embedded"`
}

func decorate(i domain.Invoice) HALDecorator {
	var links = make(map[string]Link)
	links["self"] = Link{fmt.Sprintf("/invoice/%d", i.Id)}
	for _, o := range domain.GetOperations(i) {
		if l, err := translate(o, i); err == nil {
			links[o.Name] = l
		} else {
			log.Print(err)
		}
	}
	return HALDecorator{Invoice: i, Links: links}
}

func translate(operation domain.Operation, invoice domain.Invoice) (Link, error) {
	switch operation.Name {
	case "book":
		return Link{fmt.Sprintf("/book/%d", invoice.Id)}, nil
	case "charge":
		return Link{fmt.Sprintf("/charge/%d", invoice.Id)}, nil
	case "payment":
		return Link{fmt.Sprintf("/payment/%d", invoice.Id)}, nil
	case "archive":
		return Link{fmt.Sprintf("/payment/%d", invoice.Id)}, nil
	default:
		return Link{}, errors.New(fmt.Sprintf("No translation found for operartion %s", operation.Name))
	}
}

func (j HALInvoicePresenter) Present(i interface{}) interface{} {
	var b []byte

	switch t := i.(type) {
	case []domain.Invoice:
		var result []HALDecorator
		for _, i := range t {
			result = append(result, decorate(i))
		}
		b, _ = json.Marshal(result)
	case domain.Invoice:
		decorator := decorate(t)
		decorator.Embedded = []domain.Booking{
			{Hours: 3, Description: "Aufgeräumt"},
		}
		b, _ = json.Marshal(decorator)
	}

	return b
}
