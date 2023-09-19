package version

type Strategy interface {
	StrategyBillingCreate() Response
	StrategyControlCreate() Response
}

func (ac *ApiConfig) StrategyBillingCreate() Response {
	ac.UrlPath = StrategyBillingCreate
	return ac.ApiCall()
}

func (ac *ApiConfig) StrategyControlCreate() Response {
	ac.UrlPath = StrategyControlCreate
	return ac.ApiCall()
}
