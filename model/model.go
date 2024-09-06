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
	Id string `json:"id"`
	Username string `json:"username"`
	PasswordHash string `json:"password_hash"`
}

type SaveTokenReq struct{
	UserId string `json:"user_id"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt int64 `json:"expires_at"`
}

type Error struct{
	Message string `json:"message"`
}

type CreateTaskRead struct{
	Title string `json:"title"`
}

type CreateTaskReq struct{
	UserId string `json:"user_id"`
	Title string `json:"title"`
}

type CreateTaskResp struct{
	Id string `json:"id"`
}