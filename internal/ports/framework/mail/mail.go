package mail

type MailPort interface {
	Send(to string, subject string, body string) error
	CodeSend(to string, subject string, code string) error
	WelcomeMail(to string) error
}
