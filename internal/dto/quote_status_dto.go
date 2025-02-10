package dto

type QuoteStatus string

const (
	QuoteStatusCreated  QuoteStatus = "created"
	QuoteStatusSent     QuoteStatus = "sent"
	QuoteStatusAccepted QuoteStatus = "accepted"
	QuoteStatusRejected QuoteStatus = "rejected"
)

var validQuoteStatuses = map[QuoteStatus]bool{
	QuoteStatusCreated:  true,
	QuoteStatusSent:     true,
	QuoteStatusAccepted: true,
	QuoteStatusRejected: true,
}

func (s QuoteStatus) String() string {
	return string(s)
}

// IsValid checks if the QuoteStatus is valid
func (s QuoteStatus) IsValid() bool {
	_, exists := validQuoteStatuses[s]
	return exists
}
