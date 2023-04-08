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
	fmt.Println(consts.Conf.Captcha.Length)
	driver := base64Captcha.NewDriverDigit(int(consts.Conf.Captcha.Height), int(consts.Conf.Captcha.Width),
		int(consts.Conf.Captcha.Length), float64(consts.Conf.Captcha.MaxSkew), int(consts.Conf.Captcha.DotCount))
	cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, err := cp.Generate(); err != nil {
		response.Fail(c, errors.ParamsFailed, nil)
	} else {
		response.Success(c, model.CaptchaResp{
			CaptchaID:     id,
			PicPath:       b64s,
			CaptchaLength: int(consts.Conf.Captcha.Length),
		})
	}
}
