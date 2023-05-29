package email

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net/smtp"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/jordan-wright/email"
	"golang.org/x/sync/errgroup"

	pbAny "github.com/xiaohubai/go-grpc-layout/api/any/v1"
	"github.com/xiaohubai/go-grpc-layout/configs/conf"
	"github.com/xiaohubai/go-grpc-layout/pkg/tracing"
	"github.com/xiaohubai/go-grpc-layout/pkg/utils/html"
)

type Email struct {
	Conf     *conf.Email `json:"Conf"`
	Topic    string      `json:"topic"`
	Title    string      `json:"title"`
	FilePath string      `json:"filePath"`
	HtmlText string      `json:"htmlText"`
}

func (e *Email) Send() (err error) {
	auth := smtp.PlainAuth("", e.Conf.From, e.Conf.Secret, e.Conf.Host)
	t := checkTopics(e.Topic, e.Conf.Topics)
	if t == nil {
		return errors.New("email topic not found")
	}
	cli := &email.Email{
		To:      t.To,
		From:    fmt.Sprintf("%s <%s>", e.Conf.Nickname, e.Conf.From),
		Subject: fmt.Sprintf("%s-%s(%s)", t.Subject, e.Title, e.Conf.AppName),
		HTML:    []byte(e.HtmlText),
	}

	if e.FilePath != "" {
		cli.AttachFile(e.FilePath)
	}

	address := fmt.Sprintf("%s:%d", e.Conf.Host, e.Conf.Port)
	if e.Conf.IsSsl {
		err = cli.SendWithTLS(address, auth, &tls.Config{ServerName: e.Conf.Host})
	} else {
		err = cli.Send(address, auth)
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

func SendWarn(ctx context.Context, Conf *conf.Email, title, msg string) {
	var g errgroup.Group
	g.Go(func() error {
		value := pbAny.Warn{
			DateTime: time.Now().Local().Format(time.DateTime),
			TraceID:  tracing.TraceID(ctx),
			Error:    msg,
		}

		htmlText, err := html.EmailHTMLByText("warn", value)
		if err != nil {
			return err
		}
		e := Email{
			Conf:     Conf,
			Topic:    "warn",
			Title:    title,
			HtmlText: htmlText,
		}
		return e.Send()

	})
	if err := g.Wait(); err != nil {
		log.Errorw("key", "warn", "title", "email", "msg", err.Error())
	}

	//配置是否邮件,日志双写
	if true {
		log.Errorw("key", "warn", "title", title, "msg", msg)
	}

}

func SendWarnWithFile(ctx context.Context, Conf *conf.Email, title, filePath, msg string) {
	var g errgroup.Group
	g.Go(func() error {
		value := pbAny.Warn{
			DateTime: time.Now().Local().Format(time.DateTime),
			TraceID:  tracing.TraceID(ctx),
			Error:    msg,
		}

		htmlText, err := html.EmailHTMLByText("warn", value)
		if err != nil {
			return err
		}
		e := Email{
			Conf:     Conf,
			Topic:    "warn",
			Title:    title,
			FilePath: filePath,
			HtmlText: htmlText,
		}
		return e.Send()

	})
	if err := g.Wait(); err != nil {
		log.Errorw("key", "warn", "title", "email", "msg", err.Error())
	}

	//配置是否邮件,日志双写
	if true {
		log.Errorw("key", "warn", "title", title, "msg", msg)
	}

}
