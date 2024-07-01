package user

type LoginReq struct {
	UserID   string `json:"user_id"`
	Password string `json:"password"`
}
