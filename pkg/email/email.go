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
	pbAny "github.com/xiaohubai/go-grpc-layout/api/any/v1"
	"github.com/xiaohubai/go-grpc-layout/configs/conf"
	"github.com/xiaohubai/go-grpc-layout/internal/consts"
	"github.com/xiaohubai/go-grpc-layout/pkg/tracing"
	"github.com/xiaohubai/go-grpc-layout/pkg/utils/html"
	"golang.org/x/sync/errgroup"
)

func Send(topic, title, filePath, htmlText string) (err error) {
	auth := smtp.PlainAuth("", consts.Conf.Email.From, consts.Conf.Email.Secret, consts.Conf.Email.Host)
	t := checkTopics(topic, consts.Conf.Email.Topics)
	if t == nil {
		return errors.New("email topic not found")
	}
	e := &email.Email{
		To:      t.To,
		From:    fmt.Sprintf("%s <%s>", consts.Conf.Email.Nickname, consts.Conf.Email.From),
		Subject: fmt.Sprintf("%s-%s(%s)", t.Subject, title, consts.Conf.Global.AppName),
		HTML:    []byte(htmlText),
	}

	if filePath != "" {
		e.AttachFile(filePath)
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

func SendWarn(ctx context.Context, title, msg string) {
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
		return Send("warn", title, "", htmlText)

	})
	if err := g.Wait(); err != nil {
		log.Errorw("key", "warn", "title", "email", "msg", err.Error())
	}

	//配置是否邮件,日志双写
	if true {
		log.Errorw("key", "warn", "title", title, "msg", msg)
	}

}

func SendWarnWithFile(ctx context.Context, title, filePath, msg string) {
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
		return Send("warn", title, filePath, htmlText)

	})
	if err := g.Wait(); err != nil {
		log.Errorw("key", "warn", "title", "email", "msg", err.Error())
	}

	//配置是否邮件,日志双写
	if true {
		log.Errorw("key", "warn", "title", title, "msg", msg)
	}

}
