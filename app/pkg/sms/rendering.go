package sms

import (
	"bytes"
	"html/template"
)

type templateData struct {
	From      string
	Message   string
	BucketUrl string
}

func (s *service) renderTemplate(from string, message string) (string, error) {
	// Load the template
	t, err := template.ParseFiles("assets/email.html")
	if err != nil {
		s.logger.Error("Unable to find email template")
		return "", err
	}

	// Prepare the data
	data := templateData{
		From:      from,
		Message:   message,
		BucketUrl: s.config.Hosting.BucketUrl,
	}

	// Render the template
	var buf bytes.Buffer
	err = t.Execute(&buf, data)
	if err != nil {
		s.logger.Error("Unable to render the email template")
		return "", err
	}

	return buf.String(), nil
}
