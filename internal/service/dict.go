package service

import (
	"github.com/gin-gonic/gin"

	"github.com/xiaohubai/go-grpc-layout/configs/conf"
	"github.com/xiaohubai/go-grpc-layout/internal/ecode"
	"github.com/xiaohubai/go-grpc-layout/pkg/consul"
	pconsul "github.com/xiaohubai/go-grpc-layout/pkg/consul"
	"github.com/xiaohubai/go-grpc-layout/pkg/utils/response"
)

func (s *HttpService) GetDictList(c *gin.Context) {
	dict := make(map[string]interface{}, 0)
	cli, err := pconsul.NewClient(conf.C.Consul.Host, conf.C.Consul.Token)
	if err != nil {
		response.Fail(c, ecode.GetDictListFailed, err)
		return
	}
	cul := consul.New(cli)
	_, err = cul.GetConsulKV(conf.C.Consul.Kv.DictPath, &dict)
	if err != nil {
		response.Fail(c, ecode.GetDictListFailed, err)
		return
	}
	res := map[string]interface{}{
		"filename": "dict.json",
		"dictInfo": dict,
	}
	response.Success(c, res)
}
