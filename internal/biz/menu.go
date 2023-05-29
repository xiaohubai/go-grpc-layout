package biz

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/singleflight"

	v1 "github.com/xiaohubai/go-grpc-layout/api/http/v1"
	"github.com/xiaohubai/go-grpc-layout/internal/data/model"
	"github.com/xiaohubai/go-grpc-layout/pkg/jwt"
	"github.com/xiaohubai/go-grpc-layout/pkg/utils"
)

var (
	sg singleflight.Group
)

// GetAllMenuList 获取全部路由列表
func (uc *HttpUsecase) GetAllMenuList(c *gin.Context) ([]*v1.MenuResponse, error) {
	key := "menuList"
	value, err := uc.repo.Get(c.Request.Context(), key)
	if err != nil {
		v, err, _ := sg.Do(key, func() (interface{}, error) {
			menuList, err := uc.repo.ListAllMenu(c.Request.Context())
			if err != nil {
				return nil, err
			}
			data := menuTreeHandler(menuList)
			value := utils.JsonToMarshal(data)
			uc.repo.SetEx(c.Request.Context(), key, value, 1*time.Hour)
			return value, nil
		})
		if err != nil {
			return nil, err
		}
		value = v.(string)
	}
	var res []*v1.MenuResponse
	err = json.Unmarshal([]byte(value), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetRoleMenuList 获取角色路由列表
func (uc *HttpUsecase) GetRoleMenuList(c *gin.Context) ([]*v1.MenuResponse, error) {
	claims, ok := c.Get("claims")
	if !ok {
		return nil, errors.New("token解析失败")
	}
	userInfo := claims.(*jwt.Claims)
	menuList, err := uc.repo.ListRoleMenu(c.Request.Context(), &model.Menu{RoleIDs: userInfo.RoleID})
	if err != nil {
		return nil, err
	}
	return menuTreeHandler(menuList), nil
}

func menuTreeHandler(menuList []*model.Menu) []*v1.MenuResponse {
	m := make(map[int]*v1.MenuResponse, 0)
	res := make([]*v1.MenuResponse, 0)

	for _, v := range menuList {
		data := v1.MenuResponse{
			Path:      v.Path,
			Name:      v.Name,
			Component: v.Component,
			Redirect:  v.Redirect,
			Meta: &v1.MenuResponse_Meta{
				ID:        v.ID,
				ParentID:  v.ParentID,
				RoleIDs:   v.RoleIDs,
				Title:     v.Title,
				Icon:      v.Icon,
				Hidden:    v.Hidden,
				KeepAlive: v.KeepAlive,
				Sort:      v.Sort,
			},
		}
		if v.ParentID == 0 {
			res = append(res, &data)
		} else {
			m[int(v.ParentID)].Children = append(m[int(v.ParentID)].Children, &data)
		}
		m[int(v.ID)] = &data
	}
	return res
}

func (uc *HttpUsecase) AddRoleMenu(c *gin.Context, req *v1.AddRoleMenuRequest) error {
	claims, ok := c.Get("claims")
	if !ok {
		return errors.New("token解析失败")
	}
	userInfo := claims.(*jwt.Claims)

	return uc.repo.AddRoleMenu(c.Request.Context(), &model.Menu{
		Path:       req.Path,
		Name:       req.Name,
		Component:  req.Component,
		ParentID:   req.ParentID,
		RoleIDs:    req.RoleIDs,
		Title:      req.Title,
		Icon:       req.Icon,
		Hidden:     req.Hidden,
		KeepAlive:  req.KeepAlive,
		Redirect:   req.Redirect,
		Sort:       req.Sort,
		CreateUser: userInfo.UserName,
		UpdateUser: userInfo.UserName,
	})
}

func (uc *HttpUsecase) DeleteRoleMenuByID(ctx context.Context, req *v1.DeleteRoleMenuRequest) error {
	return uc.repo.DeleteRoleMenuByID(ctx, &model.Menu{ID: req.ID})
}

func (uc *HttpUsecase) UpdateRoleMenu(c *gin.Context, req *v1.UpdateRoleMenuRequest) error {
	claims, ok := c.Get("claims")
	if !ok {
		return errors.New("token解析失败")
	}
	userInfo := claims.(*jwt.Claims)

	return uc.repo.UpdateRoleMenu(c.Request.Context(), &model.Menu{
		ID:         req.ID,
		Path:       req.Path,
		Name:       req.Name,
		Component:  req.Component,
		ParentID:   req.ParentID,
		RoleIDs:    req.RoleIDs,
		Title:      req.Title,
		Icon:       req.Icon,
		Hidden:     req.Hidden,
		KeepAlive:  req.KeepAlive,
		Redirect:   req.Redirect,
		Sort:       req.Sort,
		UpdateUser: userInfo.UserName,
	})
}
