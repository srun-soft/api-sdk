package configs

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

	// UsersVisitors 访客相关
	UsersVisitors      = "/api/v2/user/visitors"
	UsersTokenVisitors = "/api/v2/user/token-visitors"

	// UsersEventVisitors 事件访客
	UsersEventVisitors = "/api/v2/user/event-visitors"
	EventVisitorView   = "/api/v2/visitor/view-event-visitor"
	EventVisitorCreate = "/api/v2/visitor/create-event-visitor"
	EventVisitorUpdate = "/api/v2/visitor/update-event-visitor"
	EventVisitorDelete = "/api/v2/visitor/delete-event-visitor"

	// UsersViewInvite 邀请码访客
	UsersViewInvite    = "/api/v2/user/view-invite"
	UsersInviteVisitor = "/api/v2/user/invite-visitors"
	UsersUseInvite     = "/api/v2/user/use-invite"
	UsersInviteCreate  = "/api/v2/user/create-invite"
	UsersInviteDisable = "/api/v2/user/disabled-invite"

	// VisitorsWhiteCreate 访客白名单
	VisitorsWhiteCreate = "/api/v2/visitor/create-visitor-white"
	VisitorsWhiteDelete = "/api/v2/visitor/delete-visitor-white"

	// GroupsSearchAll 用户组相关
	GroupsSearchAll = "/api/v2/groups"
	GroupsSubscribe = "/api/v2/group/subscribe"
	GroupsCreate    = "/api/v2/groups"

	// StrategiesBillingCreate 计费策略相关
	StrategiesBillingCreate = ""
	StrategiesControlCreate = "/api/v2/strategy/control-create"
	ProductsView            = "/api/v2/product/view"
	ProductsCreate          = "/api/v2/product/create"
	ProductsDelete          = "/api/v2/product/delete"
	ProductsUpdate          = "/api/v2/product/update"
	ProductsIndex           = "/api/v2/product/index"

	// FinancialsRechargeWallet 财务相关
	FinancialsRechargeWallet        = "/api/v2/financial/recharge-wallet"
	FinancialsPayType               = "/api/v2/financial/pay-type"
	FinancialsCreatePayment         = "/api/v2/financial/create-payment"
	FinancialsDeletePayment         = "/api/v2/financial/delete-payment"
	FinancialsUpdatePayment         = "/api/v2/financial/update-payment"
	FinancialsExtraPay              = "/api/v2/financial/extra-pay"
	FinancialsTransfer              = "/api/v2/financial/transfer"
	FinancialsRechargeCards         = "/api/v2/financial/recharge-cards"
	FinancialsPaymentRecords        = "/api/v2/financial/payment-records"
	FinancialsPaymentOverview       = "/api/v2/financial/payment-overview"
	FinancialsPaymentDataSync       = "/api/v2/financial/payment-data-sync"
	FinancialsRefund                = "/api/v2/financial/refund"
	FinancialsPackageRechargeRecord = "/api/v2/financial/package-recharge-record"
	CheckoutListDetail              = "/api/v2/checkoutlist/detail"

	// ProductsRecharge 产品相关
	ProductsRecharge                          = "/api/v2/product/product-recharge"
	ProductsRechargeAdvanced                  = "/api/v2/product/recharge"
	ProductsUseNumber                         = "/api/v2/product/use-number"
	ProductsSubscribe                         = "/api/v2/product/subscribe"
	ProductsCanSubscribeProducts              = "/api/v2/product/can-subscribe-products"
	ProductsCancel                            = "/api/v2/product/cancel"
	ProductsRefund                            = "/api/v2/product/refund"
	ProductsExpire                            = "/api/v2/product/expire"
	ProductsDisable                           = "/api/v2/product/disable-product"
	ProductsEnable                            = "/api/v2/product/enable-product"
	ProductsNextBillingCycle                  = "/api/v2/product/next-billing-cycle"
	ProductsTransferProduct                   = "/api/v2/product/transfer-product"
	ProductsReservationTransferProducts       = "/api/v2/product/reservation-transfer-products"
	ProductsCancelReservationTransferProducts = "/api/v2/product/cancel-reservation-transfer-products"

	// PackagesUsersPackages 套餐相关
	PackagesUsersPackages = "/api/v2/package/users-packages"
	PackagesBuy           = "/api/v2/package/buy"
	PackagesBuySuper      = "/api/v2/package/buy-super"
	Packages              = "/api/v2/packages"
	PackagesBuys          = "/api/v2/package/buys"
	PackagesBatch         = "/api/v2/package/batch"

	// UsersSearchMobilePhone 运营商相关
	UsersSearchMobilePhone = "/api/v2/user/search-mobile-phone"
	ProductsCheckOperators = "/api/v2/product/check-operators"
	ProductsOperators      = "/api/v2/product/operators"

	// BasesListMacAuth 无感知相关
	BasesListMacAuth   = "/api/v2/base/list-mac-auth"
	BasesCreateMacAuth = "/api/v2/base/create-mac-auth"
	BasesDeleteMacAuth = "/api/v2/base/delete-mac-auth"
	BasesUpdateMacAuth = "/api/v2/base/update-mac-auth"
	BasesPhoneAuth     = "/api/v2/base/phone-auth"
	BasesMacs          = "/api/v2/base/macs"
	BasesCreateMac     = "/api/v2/base/create-mac"
	BasesDeleteMac     = "/api/v2/base/delete-mac"
	BasesUpdateMac     = "/api/v2/base/update-mac"
	BasesUpdateVlan    = "/api/v2/base/update-vlan"
	BasesGetIp         = "/api/v2/base/get-ip"
	BasesBindIp        = "/api/v2/base/bind-ip"

	// BasesOnlineDrop 在线相关
	BasesOnlineDrop      = "/api/v2/base/online-drop"
	BasesBatchOnlineDrop = "/api/v2/base/batch-online-drop"
	Auth                 = "/api/v2/auth"
	BasesOnlineEquipment = "/api/v2/base/online-equipment"
	BasesOnlineData      = "/api/v2/base/online-data"

	// Details 上网明细相关
	Details    = "/api/v2/details"
	DetailsTop = "/api/v2/detail/top"

	// MessagesNew 通知公告相关
	MessagesNew    = "/api/v2/message/new-message"
	Messages       = "/api/v2/message"
	MessagesNotice = "/api/v2/message/notice"

	// UsersValidateManager 管理员相关
	UsersValidateManager           = "/api/v2/user/validate-manager"
	UsersAddManager                = "/api/v2/user/add-manager"
	UsersResetPasswordManager      = "/api/v2/user/reset-password-manager"
	UsersSuperResetPasswordManager = "/api/v2/user/super-reset-password-manager"

	// LogsLogin 日志相关
	LogsLogin           = "/api/v2/log/login"
	LogsAuthErr         = "/api/v2/log/auth-err"
	LogsAuthErrorTotal  = "/api/v2/log/auth-error-total"
	LogsSys             = "/api/v2/log/sys"
	StatisticsDayByte   = "/api/v2/statistics/day-byte"
	StatisticsDayTime   = "/api/v2/statistics/day-time"
	StatisticsMonthByte = "/api/v2/statistics/month-byte"
	FlowsSearch         = "/api/v2/flow/search"
	AuditsSearch        = "/api/v2/audit/search"
)
