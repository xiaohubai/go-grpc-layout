package email

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net/smtp"
	"time"

	"github.com/jordan-wright/email"
	pbAny "github.com/xiaohubai/go-grpc-layout/api/any/v1"
	"github.com/xiaohubai/go-grpc-layout/configs"
	"github.com/xiaohubai/go-grpc-layout/internal/consts"
	"github.com/xiaohubai/go-grpc-layout/pkg/html"
	"github.com/xiaohubai/go-grpc-layout/pkg/tracing"
)

func Send(topic, htmlText string) (err error) {
	auth := smtp.PlainAuth("", consts.Cfg.Email.From, consts.Cfg.Email.Secret, consts.Cfg.Email.Host)
	t := checkTopics(topic, consts.Cfg.Email.Topics)
	if t == nil {
		return errors.New("email topic not found")
	}
	e := &email.Email{
		To:      t.To,
		From:    consts.Cfg.Email.From,
		Subject: fmt.Sprintf("%s(%s)", t.Subject, consts.Cfg.Global.AppName),
		HTML:    []byte(htmlText),
	}
	address := fmt.Sprintf("%s:%d", consts.Cfg.Email.Host, consts.Cfg.Email.Port)
	if consts.Cfg.Email.IsSsl {
		err = e.SendWithTLS(address, auth, &tls.Config{ServerName: consts.Cfg.Email.Host})
	} else {
		err = e.Send(address, auth)
	}
	return
}

func checkTopics(topic string, topics []*configs.Email_Topics) *configs.Email_Topics {
	for _, v := range topics {
		if topic == v.Name {
			return v
		}
	}
	return nil
}

func SendWarn(ctx context.Context, msg string) error {
	value := pbAny.Warn{
		DateTime: time.Now().String(),
		TraceID:  tracing.TraceID(ctx),
		Error:    msg,
	}
	htmlText, err := html.FormatText("warn", value)
	if err != nil {
		return err
	}
	return Send("warn", htmlText)
}
