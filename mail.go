package main

import (
	"bytes"
	"context"
	"text/template"
	"time"

	"github.com/Masterminds/sprig"
	"github.com/mailgun/mailgun-go/v4"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"gomodules.xyz/sets"
)

var knowTestEmails = sets.NewString("1gtm@appscode.com")
var skipEmailDomains = sets.NewString("appscode.com")

func RenderMail(src string, data interface{}) (string, string, error) {
	tpl := template.Must(template.New("").Funcs(sprig.TxtFuncMap()).Parse(src))

	var bodyText bytes.Buffer
	err := tpl.Execute(&bodyText, data)
	if err != nil {
		return "", "", err
	}

	var bodyHtml bytes.Buffer
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)
	if err := md.Convert(bodyText.Bytes(), &bodyHtml); err != nil {
		return "", "", err
	}
	return bodyText.String(), bodyHtml.String(), nil
}

type Options struct {
	// Your available domain names can be found here:
	// (https://app.mailgun.com/app/domains)
	MailgunDomain string

	// You can find the Private API Key in your Account Menu, under "Settings":
	// (https://app.mailgun.com/app/account/security)
	MailgunPrivateAPIKey string

	MailSender         string
	MailLicenseTracker string
	MailReplyTo        string
}

func SendMail(opts Options, recipient, subject, bodyText, bodyHtml string, attachments map[string][]byte) error {
	mg := mailgun.NewMailgun(opts.MailgunDomain, opts.MailgunPrivateAPIKey)

	// The message object allows you to add attachments and Bcc recipients
	msg := mg.NewMessage(opts.MailSender, subject, bodyText, recipient)
	msg.AddBCC(opts.MailLicenseTracker)
	msg.SetReplyTo(opts.MailReplyTo)

	msg.SetTracking(true)
	msg.SetTrackingClicks(true)
	msg.SetTrackingOpens(true)
	msg.SetHtml(bodyHtml)
	//msg.SetHtml("<html><body><h1>Testing some Mailgun Awesomeness!!</h1></body></html>")
	for filename, data := range attachments {
		msg.AddBufferAttachment(filename, data)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message with a 10 second timeout
	_, _, err := mg.Send(ctx, msg)
	return err
}
