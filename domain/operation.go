package domain

type Operation struct {
	Name string
}

func GetOperations(invoice *Invoice) []Operation {
	switch invoice.Status {
	case "open":
		return []Operation{{"book"}, {"charge"}}
	case "payment expected":
		return []Operation{{"payment"}}
	case "payed":
		return []Operation{{"archive"}}
	case "archived":
		return []Operation{{"revoke"}}
	case "revoked":
		return []Operation{{"archive"}}
	default:
		return []Operation{}
	}
}
