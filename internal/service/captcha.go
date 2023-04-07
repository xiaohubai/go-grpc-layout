package service

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/xiaohubai/go-grpc-layout/internal/errors"
	"github.com/xiaohubai/go-grpc-layout/internal/model"
	"github.com/xiaohubai/go-grpc-layout/pkg/configs"
	"github.com/xiaohubai/go-grpc-layout/pkg/utils/response"
)

var store = base64Captcha.DefaultMemStore

// Captcha 生成验证码
func (s *HttpService) Captcha(c *gin.Context) {
	driver := base64Captcha.NewDriverDigit(int(configs.Cfg.Captcha.Height), int(configs.Cfg.Captcha.Width),
		int(configs.Cfg.Captcha.Length), float64(configs.Cfg.Captcha.MaxSkew), int(configs.Cfg.Captcha.DotCount))
	cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, err := cp.Generate(); err != nil {
		response.Fail(c, errors.ParamsFailed, nil)
	} else {
		response.Success(c, model.CaptchaResp{
			CaptchaID:     id,
			PicPath:       b64s,
			CaptchaLength: int(configs.Cfg.Captcha.Length),
		})
	}
}
