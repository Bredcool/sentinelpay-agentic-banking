package airwallex

import (
	"context"
	"fmt"
)

type MockClient struct {
	Payments map[string]Payment
}

func NewMockClient() *MockClient {
	return &MockClient{
		Payments: map[string]Payment{
			"pay_demo_001": {
				ID:       "pay_demo_001",
				Status:   "failed",
				Currency: "USD",
				Amount:   100,
			},
		},
	}
}

func (m *MockClient) GetPayment(
	ctx context.Context,
	paymentID string,
) (Payment, error) {

	payment, ok := m.Payments[paymentID]
	if !ok {
		return Payment{}, fmt.Errorf("payment not found: %s", paymentID)
	}

	return payment, nil
}

func (m *MockClient) GetAccount(
	ctx context.Context,
	accountID string,
) (Account, error) {
	return Account{
		ID:       accountID,
		Currency: "USD",
		Balance:  1000,
	}, nil
}

func (m *MockClient) RetryPayment(
	ctx context.Context,
	paymentID string,
) error {
	return nil
}
