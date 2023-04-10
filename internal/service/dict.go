package service

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaohubai/go-grpc-layout/internal/consts"
	"github.com/xiaohubai/go-grpc-layout/internal/errors"
	"github.com/xiaohubai/go-grpc-layout/pkg/configs"
	"github.com/xiaohubai/go-grpc-layout/pkg/utils/response"
)

type Dict struct {
	List []List `json:"list"`
}

type List struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (s *HttpService) GetDictList(c *gin.Context) {
	dict := []Dict{}
	err := configs.NewConsulConfigSource(consts.Cfg.Conf.Host, consts.Cfg.Conf.DictPath, &dict)
	if err != nil {
		response.Fail(c, errors.LoginFailed, err)
		return
	}
	response.Success(c, dict)
}
