package usersvc

type SendOTPReq struct {
	Email string `json:"email"`
}

type SendOTPRes struct {
	UserPID string `json:"user_pid"`
	Email   string `json:"email"`
}

type ResendOTPReq struct {
	Email string `json:"email"`
}

type ResendOTPRes struct {
	UserPID string `json:"user_pid"`
	Email   string `json:"email"`
}

type VerifyOTPReq struct {
	UserPID string `json:"user_pid"`
	Otp     string `json:"otp"`
}

type VerifyOTPRes struct {
	UserPID      string `json:"user_pid"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type GetProfileRes struct {
	PID            string `json:"pid"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	AvailableCoins int    `json:"available_coins"`
	MobileNumber   string `json:"mobile_number"`
}
