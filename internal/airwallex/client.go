package airwallex

import "context"

type Payment struct {
	ID       string
	Status   string
	Currency string
	Amount   float64
}

type Account struct {
	ID       string
	Currency string
	Balance  float64
}

type Client interface {
	GetPayment(
		ctx context.Context,
		paymentID string,
	) (Payment, error)

	GetAccount(
		ctx context.Context,
		accountID string,
	) (Account, error)

	RetryPayment(
		ctx context.Context,
		paymentID string,
	) error
}
