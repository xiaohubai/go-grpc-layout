package ecode

const (
	Success                 = 0    // 成功
	Failed                  = 4000 // 系统内部错误
	CaptchaFailed           = 4001 // 验证码获取失败
	ParamsFailed            = 4002 // 参数校验错误
	LoginFailed             = 4003 // 登录失败
	TokenFailed             = 4004 // token无效
	TokenExpired            = 4005 // token授权已过期
	CasbinFailed            = 4006 // 权限不足
	CaptchaVerifyFailed     = 4007 // 验证码校验失败
	RegisterFailed          = 4008 // 注册失败:用户已注册
	MenuListFailed          = 4009 // 获取路由菜单失败
	CasbinAddFailed         = 4010 // 权限添加失败
	CasbinDelFailed         = 4011 // 权限删除失败
	CasbinUpdateFailed      = 4012 // 权限更新失败
	CasbinListFailed        = 4013 // 权限列表失败
	RateLimitAllowFailed    = 4014 // 超出请求频率限制
	FileWithExcelFailed     = 4015 // 文件不是excel
	FileReportFailed        = 4016 // 文件上传失败
	FileOpenFailed          = 4017 // 文件打开失败
	GetUserInfoFailed       = 4018 // 获取用户信息失败
	UpdateUserInfoFailed    = 4019 // 更新用户信息失败
	GetCasbinListFailed     = 4020 // 获取权限表信息失败
	NotAdminID              = 4021 // 无权限操作该接口
	SetCasbinFailed         = 4022 // 更新权限失败
	GetDictListFailed       = 4023 // 获取字典序失败
	GetSettingsFailed       = 4024 // 获取layout配置失败
	UpdateSettingsFailed    = 4025 // 设置layout配置失败
	TokenValidateFailed     = 4026 // token解析失败
	UpdatePasswordFailed    = 4027 // 更新用户密码失败
	RateLimitAllowErrFailed = 4028 // 请求频率限制接口报错
	DebugPerfFailed         = 4029 // 性能测试失败
)

var Msg = map[int]string{
	Success:       "成功",
	Failed:        "系统内部错误",
	CaptchaFailed: "验证码获取失败",
	ParamsFailed:  "参数校验错误",
	TokenFailed:   "token无效",
	LoginFailed:   "登录失败",

	CasbinFailed:            "权限不足",
	RegisterFailed:          "注册失败",
	CaptchaVerifyFailed:     "验证码校验失败",
	MenuListFailed:          "获取路由菜单失败",
	CasbinAddFailed:         "权限添加失败",
	CasbinDelFailed:         "权限删除失败",
	CasbinUpdateFailed:      "权限更新失败",
	CasbinListFailed:        "权限列表失败",
	RateLimitAllowFailed:    "超出请求频率限制",
	FileWithExcelFailed:     "不是excel文件",
	FileReportFailed:        "文件上传失败",
	FileOpenFailed:          "文件打开失败",
	GetUserInfoFailed:       "获取用户信息失败",
	UpdateUserInfoFailed:    "更新用户信息失败",
	GetCasbinListFailed:     "获取权限表信息失败",
	NotAdminID:              "无权限操作该接口",
	SetCasbinFailed:         "更新权限失败",
	GetDictListFailed:       "获取字典序失败",
	GetSettingsFailed:       "获取layout配置失败",
	UpdateSettingsFailed:    "设置layout配置失败",
	TokenValidateFailed:     "token解析失败",
	UpdatePasswordFailed:    "更新用户密码失败",
	RateLimitAllowErrFailed: "请求频率限制接口报错",
	DebugPerfFailed:         "性能测试失败",
}
