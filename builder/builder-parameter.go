package builder

import (
	"fmt"
	"strings"
)

func MainBuilderParameter() {
	SendEmail(func(b *EmailBuilder) {
		b.From("matt@mail.com").
			To("marry@mail.com").
			Subject("Invitation to birthday party").
			Body("Let's go party!")
	})
}

type email struct {
	from, to, subject, body string
}

type EmailBuilder struct {
	email email
}

func (b *EmailBuilder) From(from string) *EmailBuilder {
	if !strings.Contains(from, "@") {
		panic("from is not a valid email address")
	}
	b.email.from = from
	return b
}

func (b *EmailBuilder) To(to string) *EmailBuilder {
	if !strings.Contains(to, "@") {
		panic("to is not a valid email address")
	}
	b.email.to = to
	return b
}

func (b *EmailBuilder) Subject(subject string) *EmailBuilder {
	b.email.subject = subject
	return b
}

func (b *EmailBuilder) Body(body string) *EmailBuilder {
	b.email.body = body
	return b
}

func sendMailImpl(email *email) {
	fmt.Printf("Email is sent: %+v\n", *email)
}

type build func(*EmailBuilder)

func SendEmail(action build) {
	builder := EmailBuilder{}
	action(&builder)
	sendMailImpl(&builder.email)
}
