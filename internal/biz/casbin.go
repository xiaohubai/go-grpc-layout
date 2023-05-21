package biz

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"

	pbAny "github.com/xiaohubai/go-grpc-layout/api/any/v1"
	v1 "github.com/xiaohubai/go-grpc-layout/api/http/v1"
	"github.com/xiaohubai/go-grpc-layout/internal/data/model"
	"github.com/xiaohubai/go-grpc-layout/pkg/jwt"
)

// GetCasbinList 获取权限列表
func (uc *HttpUsecase) GetRoleCasbinList(c *gin.Context, req *v1.GetCasbinRequest) (*pbAny.PageResponse, error) {
	casbinList, total, err := uc.repo.ListRoleCasbin(c.Request.Context(), &model.CasbinRule{
		V0: req.RoleIDs,
		V1: req.Path,
		V2: req.Method,
	}, &v1.PageRequest{Page: req.Page, PageSize: req.PageSize})
	if err != nil {
		return nil, err
	}
	list := make([]*v1.GetCasbinResponse, 0)
	for _, v := range casbinList {
		data := &v1.GetCasbinResponse{
			ID:      v.ID,
			RoleIDs: v.V0,
			Path:    v.V1,
			Method:  v.V2,
			Desc:    v.Desc,
		}
		list = append(list, data)
	}
	res := &pbAny.PageResponse{
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    total,
		List:     list,
	}
	return res, nil
}

// AddCasbin 添加权限
func (uc *HttpUsecase) AddRoleCasbin(c *gin.Context, req *v1.AddCasbinRequest) error {
	claims, ok := c.Get("claims")
	if !ok {
		return errors.New("token解析失败")
	}
	userInfo := claims.(*jwt.Claims)
	return uc.repo.AddRoleCasbin(c.Request.Context(), &model.CasbinRule{
		Ptype:      "p",
		V0:         req.RoleIDs,
		V1:         req.Path,
		V2:         strings.ToUpper(req.Method),
		Desc:       req.Desc,
		CreateUser: userInfo.UserName,
		UpdateUser: userInfo.UserName,
	})
}

// AddCasbin 添加权限
func (uc *HttpUsecase) UpdateRoleCasbin(c *gin.Context, req *v1.UpdateCasbinRequest) error {
	claims, ok := c.Get("claims")
	if !ok {
		return errors.New("token解析失败")
	}
	userInfo := claims.(*jwt.Claims)
	return uc.repo.UpdateRoleCasbin(c.Request.Context(), &model.CasbinRule{
		ID:         req.ID,
		V0:         req.RoleIDs,
		V1:         req.Path,
		V2:         strings.ToUpper(req.Method),
		Desc:       req.Desc,
		UpdateUser: userInfo.UserName,
	})
}

// AddCasbin 添加权限
func (uc *HttpUsecase) DeleteRoleCasbin(c *gin.Context, req *v1.DeleteCasbinRequest) error {
	return uc.repo.DeleteRoleCasbin(c.Request.Context(), &model.CasbinRule{ID: req.ID})
}
