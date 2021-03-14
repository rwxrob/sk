package main

import (
	"crypto/tls"
	"fmt"
	"net/mail"
	"net/smtp"
	"strings"
)

// Email sends email through an SMTP server which must be configured in
// your SOIL config.yml file:
//
//     Email:
//       Host      : '127.0.0.1'
//       Port      : '1025'
//       From      : 'Mr. Rob <rwx@robs.io>'
//       User      : 'rwx@robs.io
//       Pass      : 'jadfoiaidifj'
//       VerifyTLS : false
//
// (That example is taken directly from a working local ProtonMail SMTP bridge setup.)
//
// Note that VerifyTLS should only be set to true when connecting to non-local
// SMTP servers and bridges (such as Gmail).
//
// The recipients can be string with angle brackets ("Mr. Rob <rob@rwx.io>").
// Each recipient will be send a copy of the message but will not be able to
// see other recipients listed (BCC). This allows sending the same message to
// many without initiating a separate session for each individual recipient.
// The subject should be a string with no line returns.  The body can be
// a simple string with multiple lines or it can be a complicated multipart
// message with HTML and other attachments.
func Email(recipients []string, subject, body string) error {
	cfg := LoadConfig()

	from, err := mail.ParseAddress(cfg.Email.From)
	if err != nil {
		return err
	}

	c, err := smtp.Dial(cfg.Email.Host + ":" + cfg.Email.Port)
	if err != nil {
		return err
	}

	tls := new(tls.Config)
	tls.ServerName = cfg.Email.Host
	tls.InsecureSkipVerify = !cfg.Email.VerifyTLS

	c.StartTLS(tls)

	auth := smtp.PlainAuth("", cfg.Email.User, cfg.Email.Pass, cfg.Email.Host)

	if err = c.Auth(auth); err != nil {
		return err
	}

	if err = c.Mail(from.Address); err != nil {
		return err
	}

	for _, rcpt := range recipients {
		to, err := mail.ParseAddress(rcpt)
		if err = c.Rcpt(to.Address); err != nil {
			return err
		}
	}

	w, err := c.Data()
	if err != nil {
		return err
	}

	// omit "To:" to force bcc recipients
	message := fmt.Sprintf("To: %v\r\nFrom: %v\r\nSubject: %v\r\n\r\n%v",
		strings.Join(recipients, ","), from, subject, body)

	_, err = w.Write([]byte(message))
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	c.Quit()
	return nil
}
