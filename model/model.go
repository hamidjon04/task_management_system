package model

type RegisterReq struct{
	Email string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterResp struct{
	UserId string `json:"user_id"`
}

type LoginReq struct{
	Email string `json:"email"`
	Password string `json:"password"`
}

type LoginResp struct{
	RefreshToken string `json:"refresh_token"`
	AcessToken string `json:"access_token"`
}

type UserInfo struct{
	UserId string `json:"user_id"`
	Username string `json:"username"`
	PasswordHash string `json:"password_hash"`
}

type SaveTokenReq struct{
	Userid string `json:"user_id"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt string `json:"expires_at"`
}
