package subscriptionplansvc

type GetAllSubscriptionPlansRes []struct {
	PID   string  `json:"pid"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Coins int     `json:"coins"`
}
