package mailer

import (
	"bytes"
	"html/template"
	"mail-service/models"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"
	"gopkg.in/gomail.v2"
)

type MailerInterface interface {
	SendSMTPMessage(models.Message) error
	SendGoMail(models.Message) error
}

type mailer struct {
	mail models.Mail
}

func NewMailer(m models.Mail) MailerInterface {
	return &mailer{m}
}

func (m *mailer) SendSMTPMessage(msg models.Message) error {
	if msg.From == "" {
		msg.From = m.mail.FromAddress
	}
	if msg.FromName == "" {
		msg.FromName = m.mail.FromName
	}

	data := map[string]any{
		"message": msg.Data,
	}

	msg.DataMap = data

	formattedMessage, err := m.buildHTMLMessage(msg)
	if err != nil {
		return err
	}

	server := mail.NewSMTPClient()
	server.Host = m.mail.Host
	server.Port = m.mail.Port
	server.Username = m.mail.Username
	server.Password = m.mail.Password
	server.Encryption = m.getEncryption(m.mail.Encryption)
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	smtpClient, err := server.Connect()
	if err != nil {
		return err
	}

	email := mail.NewMSG()
	email.SetFrom(msg.From).AddTo(msg.To).SetSubject(msg.Subject)
	email.AddAlternative(mail.TextHTML, formattedMessage)

	if len(msg.Attachments) > 0 {
		for _, v := range msg.Attachments {
			email.AddAttachment(v)
		}
	}

	err = email.Send(smtpClient)
	if err != nil {
		return err
	}
	return nil
}

func (m *mailer) SendGoMail(mailData models.Message) error {
	body, _ := m.buildHTMLMessage(mailData)

	msg := gomail.NewMessage()
	msg.SetHeader("From", mailData.From)
	msg.SetHeader("To", mailData.To)
	msg.SetHeader("Subject", mailData.Subject)
	msg.SetBody("text/html", body)
	// msg.Attach("/home/User/cat.jpg")

	n := gomail.NewDialer(m.mail.Host, m.mail.Port, m.mail.Username, m.mail.Password)

	// Send the email
	if err := n.DialAndSend(msg); err != nil {
		return err
	}
	return nil
}

func (m *mailer) getEncryption(s string) mail.Encryption {
	switch s {
	case "tls":
		return mail.EncryptionSTARTTLS
	case "ssl":
		return mail.EncryptionSSLTLS
	case "none":
		return mail.EncryptionNone
	default:
		return mail.EncryptionSTARTTLS
	}
}

func (m *mailer) buildHTMLMessage(msg models.Message) (string, error) {
	templateToRender := "./templates/mail.html"

	t, err := template.New("email-html").ParseFiles(templateToRender)
	if err != nil {
		return "", err
	}
	var tpl bytes.Buffer
	if err = t.ExecuteTemplate(&tpl, "body", msg.DataMap); err != nil {
		return "", err
	}

	formattedMessage := tpl.String()
	return formattedMessage, nil
}
