package main

import "fmt"

type PaymentStrategy interface {
	Pay()
}

// Credit Card Strategy Type
type CreditCardStrategy struct{}

func (c *CreditCardStrategy) Pay() {
	fmt.Println("Paying using Credit Card")
}

// Debit Card Strategy Type
type DebitCardStrategy struct{}

func (d *DebitCardStrategy) Pay() {
	fmt.Println("Paying using Debit Card")
}

// Visa Card Strategy Type
type VisaCardStrategy struct{}

func (v *VisaCardStrategy) Pay() {
	fmt.Println("Paying using Visa Card")
}

// This type sets the strategy dynamically
type PaymentMethod struct {
	paymentStrategy PaymentStrategy
}

func (p *PaymentMethod) setPaymentMethodStrategy(strategy PaymentStrategy) {
	p.paymentStrategy = strategy
}

func (p *PaymentMethod) checkOut() {
	p.paymentStrategy.Pay()
}

func main() {
	paymentMethod := &PaymentMethod{}

	// Credit Card
	creditCardStrategy := &CreditCardStrategy{}
	paymentMethod.setPaymentMethodStrategy(creditCardStrategy)
	paymentMethod.checkOut()

	// Debit Card
	debitCardStrategy := &DebitCardStrategy{}
	paymentMethod.setPaymentMethodStrategy(debitCardStrategy)
	paymentMethod.checkOut()

	// Visa Card
	visaCardStrategy := &VisaCardStrategy{}
	paymentMethod.setPaymentMethodStrategy(visaCardStrategy)
	paymentMethod.checkOut()
}
