package util

import (
	"github.com/jordan-wright/email"
	"go.uber.org/zap"
	"net/smtp"
	"net/textproto"
	"p-cy/config"
	"p-cy/log"
)

func SendAttachEmail(data []byte, dataType int, filePath string, to ...string) error {
	e := email.Email{
		To:      to,
		From:    config.GlobalConfig.Email.From,
		Subject: config.GlobalConfig.Email.Subject,
		Cc:      []string{},
		Bcc:     []string{},
		Headers: textproto.MIMEHeader{},
	}
	if dataType == 1 {
		e.Text = data
	} else {
		e.HTML = data
	}
	_, err := e.AttachFile(filePath)
	if err != nil {
		log.Logger.Error("Email attach load failed", zap.Error(err))
	}
	return e.Send(config.GlobalConfig.Email.Host+":"+config.GlobalConfig.Email.Port, smtp.PlainAuth("", config.GlobalConfig.Email.From, config.GlobalConfig.Email.Password, config.GlobalConfig.Email.Password))
}

func SendEmail(data []byte, dataType int, to ...string) error {
	e := email.Email{
		To:      to,
		From:    config.GlobalConfig.Email.From,
		Subject: config.GlobalConfig.Email.Subject,
		Cc:      []string{},
		Bcc:     []string{},
		Headers: textproto.MIMEHeader{},
	}
	if dataType == 1 {
		e.Text = data
	} else {
		e.HTML = data
	}
	return e.Send(config.GlobalConfig.Email.Host+":"+config.GlobalConfig.Email.Port, smtp.PlainAuth("", config.GlobalConfig.Email.From, config.GlobalConfig.Email.Password, config.GlobalConfig.Email.Password))
}
