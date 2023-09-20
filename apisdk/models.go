package apisdk

const (
	KeyApiSdkAccessToken = "key:api:sdk:access_token"
	AuthGetAccessToken   = "/api/v2/auth/get-access-token"

	// Users

	Users      = "/api/v2/users"
	UserUpdate = "/api/v2/user/update"
	UserDelete = "/api/v2/user/delete"

	UserCode                = "/api/v2/user/code"
	UserForgetResetPassword = "/api/v2/user/forget-reset-password"
	AuthCheckModifyPassword = "/api/v2/auth/check-modify-password"
	UserGetPassword         = "/api/v2/user/get-password"
	UserSuperResetPassword  = "/api/v2/user/super-reset-password"
	UserResetPassword       = "/api/v2/user/reset-password"

	UserView        = "/api/v2/user/view"
	UserSuperSearch = "/api/v2/user/super-search"
	UserSearch      = "/api/v2/user/search"

	UserUserStatusControl = "/api/v2/user/user-status-control"
	UserValidateUsers     = "/api/v2/user/validate-users"
	UserBalance           = "/api/v2/user/balance"
	UserSendCode          = "/api/v2/user/send-code"
	UserMaxOnlineNum      = "/api/v2/user/max-online-num"
	UserBindingPhone      = "/api/v2/user/binding-phone"

	// UserVisitors 访客相关
	UserVisitors      = "/api/v2/user/visitors"
	UserTokenVisitors = "/api/v2/user/token-visitors"

	// UserEventVisitors 事件访客
	UserEventVisitors         = "/api/v2/user/event-visitors"
	VisitorViewEventVisitor   = "/api/v2/visitor/view-event-visitor"
	VisitorCreateEventVisitor = "/api/v2/visitor/create-event-visitor"
	VisitorUpdateEventVisitor = "/api/v2/visitor/update-event-visitor"
	VisitorDeleteEventVisitor = "/api/v2/visitor/delete-event-visitor"

	// UserViewInvite 邀请码访客
	UserViewInvite     = "/api/v2/user/view-invite"
	UserInviteVisitors = "/api/v2/user/invite-visitors"
	UserUseInvite      = "/api/v2/user/use-invite"
	UserCreateInvite   = "/api/v2/user/create-invite"
	UserDisableInvite  = "/api/v2/user/disabled-invite"

	// VisitorCreateVisitorWhite 访客白名单
	VisitorCreateVisitorWhite = "/api/v2/visitor/create-visitor-white"
	VisitorDeleteVisitorWhite = "/api/v2/visitor/delete-visitor-white"

	// Groups 用户组相关
	Groups         = "/api/v2/groups"
	GroupSubscribe = "/api/v2/group/subscribe"

	// StrategyBillingCreate 计费策略相关
	StrategyBillingCreate = "/api/v2/strategy/billing-create"
	StrategyControlCreate = "/api/v2/strategy/control-create"
	ProductView           = "/api/v2/product/view"
	ProductCreate         = "/api/v2/product/create"
	ProductDelete         = "/api/v2/product/delete"
	ProductUpdate         = "/api/v2/product/update"
	ProductIndex          = "/api/v2/product/index"

	// FinancialRechargeWallet 财务相关
	FinancialRechargeWallet        = "/api/v2/financial/recharge-wallet"
	FinancialPayType               = "/api/v2/financial/pay-type"
	FinancialCreatePayment         = "/api/v2/financial/create-payment"
	FinancialDeletePayment         = "/api/v2/financial/delete-payment"
	FinancialUpdatePayment         = "/api/v2/financial/update-payment"
	FinancialExtraPay              = "/api/v2/financial/extra-pay"
	FinancialTransfer              = "/api/v2/financial/transfer"
	FinancialRechargeCards         = "/api/v2/financial/recharge-cards"
	FinancialPaymentRecords        = "/api/v2/financial/payment-records"
	FinancialPaymentOverview       = "/api/v2/financial/payment-overview"
	FinancialPaymentDataSync       = "/api/v2/financial/payment-data-sync"
	FinancialRefund                = "/api/v2/financial/refund"
	FinancialPackageRechargeRecord = "/api/v2/financial/package-recharge-record"
	CheckoutListDetail             = "/api/v2/checkoutlist/detail"

	// ProductProductRecharge 产品相关
	ProductProductRecharge                   = "/api/v2/product/product-recharge"
	ProductRecharge                          = "/api/v2/product/recharge"
	ProductUseNumber                         = "/api/v2/product/use-number"
	ProductSubscribe                         = "/api/v2/product/subscribe"
	ProductCanSubscribeProducts              = "/api/v2/product/can-subscribe-products"
	ProductCancel                            = "/api/v2/product/cancel"
	ProductRefund                            = "/api/v2/product/refund"
	ProductExpire                            = "/api/v2/product/expire"
	ProductDisableProduct                    = "/api/v2/product/disable-product"
	ProductEnableProduct                     = "/api/v2/product/enable-product"
	ProductNextBillingCycle                  = "/api/v2/product/next-billing-cycle"
	ProductTransferProduct                   = "/api/v2/product/transfer-product"
	ProductReservationTransferProducts       = "/api/v2/product/reservation-transfer-products"
	ProductCancelReservationTransferProducts = "/api/v2/product/cancel-reservation-transfer-products"

	// PackageUsersPackages 套餐相关
	PackageUsersPackages = "/api/v2/package/users-packages"
	PackageBuy           = "/api/v2/package/buy"
	PackageBuySuper      = "/api/v2/package/buy-super"
	Packages             = "/api/v2/packages"
	PackageBuys          = "/api/v2/package/buys"
	PackageBatch         = "/api/v2/package/batch"

	// UserSearchMobilePhone 运营商相关
	UserSearchMobilePhone = "/api/v2/user/search-mobile-phone"
	ProductCheckOperators = "/api/v2/product/check-operators"
	ProductOperators      = "/api/v2/product/operators"

	// BaseListMacAuth 无感知相关
	BaseListMacAuth   = "/api/v2/base/list-mac-auth"
	BaseCreateMacAuth = "/api/v2/base/create-mac-auth"
	BaseDeleteMacAuth = "/api/v2/base/delete-mac-auth"
	BaseUpdateMacAuth = "/api/v2/base/update-mac-auth"
	BasePhoneAuth     = "/api/v2/base/phone-auth"
	BaseMacs          = "/api/v2/base/macs"
	BaseCreateMac     = "/api/v2/base/create-mac"
	BaseDeleteMac     = "/api/v2/base/delete-mac"
	BaseUpdateMac     = "/api/v2/base/update-mac"
	BaseUpdateVlan    = "/api/v2/base/update-vlan"
	BaseGetIp         = "/api/v2/base/get-ip"
	BaseBindIp        = "/api/v2/base/bind-ip"

	// BaseOnlineDrop 在线相关
	BaseOnlineDrop      = "/api/v2/base/online-drop"
	BaseBatchOnlineDrop = "/api/v2/base/batch-online-drop"
	Auth                = "/api/v2/auth"
	BaseOnlineEquipment = "/api/v2/base/online-equipment"
	BaseOnlineData      = "/api/v2/base/online-data"

	// Details 上网明细相关
	Details   = "/api/v2/details"
	DetailTop = "/api/v2/detail/top"

	// MessageNew 通知公告相关
	MessageNew    = "/api/v2/message/new-message"
	Message       = "/api/v2/message"
	MessageNotice = "/api/v2/message/notice"

	// UserValidateManager 管理员相关
	UserValidateManager           = "/api/v2/user/validate-manager"
	UserAddManager                = "/api/v2/user/add-manager"
	UserResetPasswordManager      = "/api/v2/user/reset-password-manager"
	UserSuperResetPasswordManager = "/api/v2/user/super-reset-password-manager"

	// LogLogin 日志相关
	LogLogin            = "/api/v2/log/login"
	LogAuthErr          = "/api/v2/log/auth-err"
	LogAuthErrorTotal   = "/api/v2/log/auth-error-total"
	LogSys              = "/api/v2/log/sys"
	StatisticsDayByte   = "/api/v2/statistics/day-byte"
	StatisticsDayTime   = "/api/v2/statistics/day-time"
	StatisticsMonthByte = "/api/v2/statistics/month-byte"
	FlowSearch          = "/api/v2/flow/search"
	AuditSearch         = "/api/v2/audit/search"
)
