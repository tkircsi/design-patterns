package adapter

import (
	"fmt"
	"math/rand"
	"time"
)

func MainBankPaypal() {
	johnAccount := &Account{
		Email:  "john@mail.hu",
		Method: &BankAdapter{},
	}

	janeAccount := &Account{
		Email:  "jane@mail.hu",
		Method: &PaypalAdapter{},
	}

	johnAccount.Pay(janeAccount, 100)

	janeAccount.Pay(johnAccount, 500)
}

type PaymentMethod interface {
	SendMoney(fromEmail, toEmail string, amount int) error
}

type Account struct {
	Email  string
	Method PaymentMethod
}

func (a *Account) Pay(receiver *Account, amount int) error {
	return a.Method.SendMoney(a.Email, receiver.Email, amount)
}

// Bank
func ProcessPayment(fromAccount, toAccount, amount int) error {
	fmt.Printf("Transfered $%d  from %d to %d at %s via bank transfer\n", amount, fromAccount, toAccount, time.Now().Format("2006.01.02 15:04:05 -0700 MST"))
	return nil
}

// Bank adapter
type BankAdapter struct{}

func (a *BankAdapter) SendMoney(fromEmail, toEmail string, amount int) error {
	fromAccount := findAccountByEmail(fromEmail)
	toAccount := findAccountByEmail(toEmail)

	return ProcessPayment(fromAccount, toAccount, amount)
}

func findAccountByEmail(email string) int {
	return rand.Int()
}

// Paypal
func Send(senderMobile, receiverMobile string, amount int) error {
	fmt.Printf("Send $%d  from %s to %s at %s via paypal transfer\n", amount, senderMobile, receiverMobile, time.Now().Format("2006.01.02 15:04:05 -0700 MST"))
	return nil
}

// Paypal adapter
type PaypalAdapter struct{}

func (a *PaypalAdapter) SendMoney(fromEmail, toEmail string, amount int) error {
	fromMobile := findMobileByEmail(fromEmail)
	toMobile := findMobileByEmail(toEmail)
	return Send(fromMobile, toMobile, amount)
}

func findMobileByEmail(email string) string {
	return "36-30-233-4455"
}
