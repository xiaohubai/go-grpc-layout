package service

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaohubai/go-grpc-layout/internal/consts"
	"github.com/xiaohubai/go-grpc-layout/internal/errors"
	"github.com/xiaohubai/go-grpc-layout/pkg/consul"
	"github.com/xiaohubai/go-grpc-layout/pkg/utils/response"
)

func (s *HttpService) GetDictList(c *gin.Context) {
	dict := make(map[string]interface{}, 0)
	_, err := consul.GetConsulKV(consts.Cfg.Consul.Kv.DictPath, &dict)
	if err != nil {
		response.Fail(c, errors.GetDictListFailed, err)
		return
	}
	res := map[string]interface{}{
		"filename": "dict.json",
		"dictInfo": dict,
	}
	response.Success(c, res)
}
