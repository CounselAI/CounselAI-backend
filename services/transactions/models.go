package transactions

type CreateTransactionReq struct {
	SubscriptionPlanPID string `json:"subscription_plan_pid" binding:"required"`
}
