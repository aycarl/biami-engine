package service

import "fmt"

// PayoutProvider defines the contract for different payment gateways (Stripe, Mobile Money)
type PayoutProvider interface {
	SendPayment(amount float64, recipient string) error
}

// MockPayout handles payments during development
type MockPayout struct{}

func (m *MockPayout) SendPayment(amount float64, recipient string) error {
	fmt.Printf("Architect Log: Payment of $%.2f sent to %s via Mock Provider\n", amount, recipient)
	return nil
}
