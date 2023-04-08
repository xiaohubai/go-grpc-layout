package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/xiaohubai/go-grpc-layout/internal/consts"
	"github.com/xiaohubai/go-grpc-layout/internal/errors"
	"github.com/xiaohubai/go-grpc-layout/internal/model"
	"github.com/xiaohubai/go-grpc-layout/pkg/utils/response"
)

var store = base64Captcha.DefaultMemStore

// Captcha 生成验证码
func (s *HttpService) Captcha(c *gin.Context) {
	fmt.Println(consts.Cfg.Captcha.Length)
	driver := base64Captcha.NewDriverDigit(int(consts.Cfg.Captcha.Height), int(consts.Cfg.Captcha.Width),
		int(consts.Cfg.Captcha.Length), float64(consts.Cfg.Captcha.MaxSkew), int(consts.Cfg.Captcha.DotCount))
	cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, err := cp.Generate(); err != nil {
		response.Fail(c, errors.ParamsFailed, nil)
	} else {
		response.Success(c, model.CaptchaResp{
			CaptchaID:     id,
			PicPath:       b64s,
			CaptchaLength: int(consts.Cfg.Captcha.Length),
		})
	}
}
