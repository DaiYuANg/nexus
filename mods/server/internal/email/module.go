package email

import (
	"github.com/domodwyer/mailyak/v3"
	"go.uber.org/fx"
	"net/smtp"
)

var Module = fx.Module("email", fx.Provide(newEmailClient))

func newEmailClient() *mailyak.MailYak {
	return mailyak.New("mail.host.com:25", smtp.PlainAuth("", "user", "pass", "mail.host.com"))
}
