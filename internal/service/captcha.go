package service

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	v1 "github.com/xiaohubai/go-grpc-layout/api/http/v1"
	"github.com/xiaohubai/go-grpc-layout/internal/consts"
	"github.com/xiaohubai/go-grpc-layout/internal/errors"
	"github.com/xiaohubai/go-grpc-layout/pkg/utils/response"
)

var store = base64Captcha.DefaultMemStore

// Captcha 生成验证码
func (s *HttpService) Captcha(c *gin.Context) {
	driver := base64Captcha.NewDriverDigit(int(consts.Conf.Captcha.Height), int(consts.Conf.Captcha.Width),
		int(consts.Conf.Captcha.Length), float64(consts.Conf.Captcha.MaxSkew), int(consts.Conf.Captcha.DotCount))
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := cp.Generate()
	if err != nil {
		response.Fail(c, errors.CaptchaFailed, nil)
		return
	}
	response.Success(c, v1.CaptchaResponse{
		CaptchaID:     id,
		PicPath:       b64s,
		CaptchaLength: consts.Conf.Captcha.Length,
	})
}
