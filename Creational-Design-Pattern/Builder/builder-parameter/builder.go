package builderparameter

import "strings"

type email struct {
	from, to, subject, body string
}

type EmailBuilder struct {
	email email
}

func (eb *EmailBuilder) From(from string) *EmailBuilder {
	if !strings.Contains(from, "@") {
		panic("from is not a valid email address")
	}
	eb.email.from = from
	return eb
}

func (eb *EmailBuilder) To(to string) *EmailBuilder {
	if !strings.Contains(to, "@") {
		panic("to is not a valid email address")
	}
	eb.email.to = to
	return eb
}

func (eb *EmailBuilder) Subject(subject string) *EmailBuilder {
	eb.email.subject = subject
	return eb
}

func (eb *EmailBuilder) Body(body string) *EmailBuilder {
	eb.email.body = body
	return eb
}

func (eb *EmailBuilder) Build() email {
	return eb.email
}

type build func(*EmailBuilder)

func SendEmail(action build) email {
	eb := EmailBuilder{}
	action(&eb)
	return eb.Build()
}
