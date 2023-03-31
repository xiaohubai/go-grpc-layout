package model

type User struct {
	ID       int32  `json:"id"`
	UID      string `json:"uid"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	NickName string `json:"nickName"`
	Birth    string `json:"birth"`
	Avatar   string `json:"avatar"`
	RoleID   string `json:"roleID"`
	RoleName string `json:"roleName"`
	Phone    string `json:"phone"`
	Wechat   string `json:"wechat"`
	Email    string `json:"email"`
	State    string `json:"state"`
	Motto    string `json:"motto"`
}
