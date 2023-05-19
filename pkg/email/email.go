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
	"github.com/xiaohubai/go-grpc-layout/configs/conf"
	"github.com/xiaohubai/go-grpc-layout/internal/consts"
	"github.com/xiaohubai/go-grpc-layout/pkg/html"
	"github.com/xiaohubai/go-grpc-layout/pkg/tracing"
)

func Send(topic, htmlText string) (err error) {
	auth := smtp.PlainAuth("", consts.Conf.Email.From, consts.Conf.Email.Secret, consts.Conf.Email.Host)
	t := checkTopics(topic, consts.Conf.Email.Topics)
	if t == nil {
		return errors.New("email topic not found")
	}
	e := &email.Email{
		To:      t.To,
		From:    consts.Conf.Email.From,
		Subject: fmt.Sprintf("%s(%s)", t.Subject, consts.Conf.Global.AppName),
		HTML:    []byte(htmlText),
	}
	address := fmt.Sprintf("%s:%d", consts.Conf.Email.Host, consts.Conf.Email.Port)
	if consts.Conf.Email.IsSsl {
		err = e.SendWithTLS(address, auth, &tls.Config{ServerName: consts.Conf.Email.Host})
	} else {
		err = e.Send(address, auth)
	}
	return
}

func checkTopics(topic string, topics []*conf.Email_Topics) *conf.Email_Topics {
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
