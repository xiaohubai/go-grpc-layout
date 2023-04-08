package service

import (
	"github.com/gin-gonic/gin"
)

type Dict struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (s *HttpService) GetMenuList(c *gin.Context) {
	/* dict := &Dict{}
	configs.NewConsulConfigSource(consts.Cfg.Conf.Host, consts.Cfg.Conf.DictPath, dict)
	if err != nil {
		response.Fail(c, errors.LoginFailed, err)
		return
	}
	response.Success(c, data) */
}
