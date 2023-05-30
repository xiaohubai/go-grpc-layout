package service

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"

	v1 "github.com/xiaohubai/go-grpc-layout/api/http/v1"
	"github.com/xiaohubai/go-grpc-layout/configs/conf"
	"github.com/xiaohubai/go-grpc-layout/internal/ecode"
	"github.com/xiaohubai/go-grpc-layout/pkg/utils/response"
)

var store = base64Captcha.DefaultMemStore

// Captcha 生成验证码
func (s *HttpService) Captcha(c *gin.Context) {
	driver := base64Captcha.NewDriverDigit(int(conf.C.Captcha.Height), int(conf.C.Captcha.Width),
		int(conf.C.Captcha.Length), float64(conf.C.Captcha.MaxSkew), int(conf.C.Captcha.DotCount))
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := cp.Generate()
	if err != nil {
		response.Fail(c, ecode.CaptchaFailed, nil)
		return
	}
	response.Success(c, v1.CaptchaResponse{
		CaptchaID:     id,
		PicPath:       b64s,
		CaptchaLength: conf.C.Captcha.Length,
	})
}
