package bank

type Money int64

type Currency string

const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

type PAN string

type Card struct {
	ID       int
	PAN      string
	Balance  Money
	Currency Currency
	Color    string
	Name     string
	Active   bool
}
