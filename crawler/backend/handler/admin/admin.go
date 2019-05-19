package admin

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
	RoleId   int64  `json:"role_id"`
	Name    string `json:"name"`
}

