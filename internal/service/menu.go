package service

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/xiaohubai/go-grpc-layout/api/http/v1"
	"github.com/xiaohubai/go-grpc-layout/internal/errors"
	"github.com/xiaohubai/go-grpc-layout/pkg/utils/request"
	"github.com/xiaohubai/go-grpc-layout/pkg/utils/response"
)

func (s *HttpService) GetAllMenuList(c *gin.Context) {
	req := &v1.PageRequest{}
	err := request.ShouldBindJSON(c, req)
	if err != nil {
		response.Fail(c, errors.ParamsFailed, err)
		return
	}

	data, err := s.uc.GetAllMenuList(c.Request.Context(), req)
	if err != nil {
		response.Fail(c, errors.MenuListFailed, err)
		return
	}
	response.Success(c, data)
}

func (s *HttpService) GetRoleMenuList(c *gin.Context) {

}
