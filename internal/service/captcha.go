package service

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	v1 "github.com/xiaohubai/go-grpc-layout/api/errors/v1"
	"github.com/xiaohubai/go-grpc-layout/pkg/configs"
	"github.com/xiaohubai/go-grpc-layout/pkg/utils/response"
)

type CaptchaResp struct {
	CaptchaID     string `json:"captchaId"`
	PicPath       string `json:"picPath"`
	CaptchaLength int    `json:"captchaLength"`
}

var store = base64Captcha.DefaultMemStore

// Captcha 生成验证码
func (s *HttpService) Captcha(c *gin.Context) {
	driver := base64Captcha.NewDriverDigit(int(configs.Cfg.Captcha.Height), int(configs.Cfg.Captcha.Width),
		int(configs.Cfg.Captcha.Length), float64(configs.Cfg.Captcha.MaxSkew), int(configs.Cfg.Captcha.DotCount))
	cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, err := cp.Generate(); err != nil {
		response.Fail(c, v1.Error_RequestFail, nil)
	} else {
		response.Success(c, CaptchaResp{
			CaptchaID:     id,
			PicPath:       b64s,
			CaptchaLength: int(configs.Cfg.Captcha.Length),
		})
	}
}
