package sdk

import (
	"github.com/srun-soft/api-sdk/configs"
	"github.com/srun-soft/api-sdk/tools"
)

// 策略相关

// StrategyBillingCreate 添加计费策略
func (APIClient) StrategyBillingCreate(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.StrategiesBillingCreate, data)
}

// StrategyControlCreate 添加控制策略
func (APIClient) StrategyControlCreate(data map[string]interface{}) (tools.SrunResponse, error) {
	return ApiCall(configs.StrategiesControlCreate, data)
}
