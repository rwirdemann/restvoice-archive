package rest

import (
	"encoding/json"
	"fmt"
	"github.com/rwirdemann/restvoice/domain"
	"errors"
	"log"
)

type RVInvoicePresenter struct {
}

func NewRVInvoicePresenter() RVInvoicePresenter {
	return RVInvoicePresenter{}
}

func (j RVInvoicePresenter) present(i interface{}) interface{} {
	invoice := i.(*domain.Invoice)
	b, _ := json.Marshal(decorate(invoice))
	return b
}

type Link struct {
	Method string `json:"method"`
	Href   string `json:"href"`
}

type LinksDecorator struct {
	*domain.Invoice
	Links map[string]Link `json:"_links"`
}

func decorate(i *domain.Invoice) LinksDecorator {
	var links = make(map[string]Link)
	links["self"] = Link{"GET", fmt.Sprintf("/invoice/%d", i.Id)}
	for _, o := range domain.GetOperations(i) {
		if l, err := translate(o, i); err == nil {
			links[o.Name] = l
		} else {
			log.Print(err)
		}
	}
	return LinksDecorator{Invoice: i, Links: links}
}

func translate(operation domain.Operation, invoice *domain.Invoice) (Link, error) {
	switch operation.Name {
	case "book":
		return Link{"POST", fmt.Sprintf("/book/%d", invoice.Id)}, nil
	case "charge":
		return Link{"PUT", fmt.Sprintf("/charge/%d", invoice.Id)}, nil
	case "payment":
		return Link{"PUT", fmt.Sprintf("/payment/%d", invoice.Id)}, nil
	case "archive":
		return Link{"DELETE", fmt.Sprintf("/payment/%d", invoice.Id)}, nil
	default:
		return Link{}, errors.New(fmt.Sprintf("No translation found for operartion %s", operation.Name))
	}
}

func (j RVInvoicePresenter) Present(i interface{}) interface{} {
	var b []byte

	switch t := i.(type) {
	case []*domain.Invoice:
		var result []LinksDecorator
		for _, i := range t {
			result = append(result, decorate(i))
		}
		b, _ = json.Marshal(result)
	case *domain.Invoice:
		b, _ = json.Marshal(decorate(t))
	}

	return b
}
