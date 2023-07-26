package sdk

const (
	// UsersCreate 用户基本操作相关
	UsersCreate = "/api/v2/users"
	UsersUpdate = "/api/v2/user/update"
	UsersDelete = "/api/v2/user/delete"

	// UsersCode 密码相关
	UsersCode                = "/api/v2/user/code"
	UsersForgetResetPassword = "/api/v2/user/forget-reset-password"
	AuthCheckModifyPassword  = "/api/v2/auth/check-modify-password"
	UsersGetPassword         = "/api/v2/user/get-password"
	UsersSuperResetPassword  = "/api/v2/user/super-reset-password"
	UsersResetPassword       = "/api/v2/user/reset-password"

	// UsersView 查询用户相关
	UsersView        = "/api/v2/user/view"
	UsersSuperSearch = "/api/v2/user/super-search"
	UsersSearch      = "/api/v2/user/search"

	// UsersStatusControl 其他辅助接口
	UsersStatusControl = "/api/v2/user/user-status-control"
	UsersValidate      = "/api/v2/user/validate-users"
	UsersBalance       = "/api/v2/user/balance"
	UsersSendCode      = "/api/v2/user/send-code"
	UsersMaxOnlineNum  = "/api/v2/user/max-online-num"
	UsersBindingPhone  = "/api/v2/user/binding-phone"
)
