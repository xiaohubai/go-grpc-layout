package service

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaohubai/go-grpc-layout/internal/consts"
	"github.com/xiaohubai/go-grpc-layout/internal/errors"
	"github.com/xiaohubai/go-grpc-layout/pkg/configs"
	"github.com/xiaohubai/go-grpc-layout/pkg/utils/response"
)

func (s *HttpService) GetSettings(c *gin.Context) {
	dict := make(map[string]interface{}, 0)
	err := configs.NewConsulConfigSource(consts.Cfg.Conf.Host, consts.Cfg.Conf.DictPath, &dict)
	if err != nil {
		response.Fail(c, errors.LoginFailed, err)
		return
	}
	response.Success(c, dict)
}

func (s *HttpService) SetSettings(c *gin.Context) {
	dict := make(map[string]interface{}, 0)
	err := configs.NewConsulConfigSource(consts.Cfg.Conf.Host, consts.Cfg.Conf.DictPath, &dict)
	if err != nil {
		response.Fail(c, errors.LoginFailed, err)
		return
	}
	response.Success(c, dict)
}
