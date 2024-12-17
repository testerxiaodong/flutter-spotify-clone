package types

type SignupReq struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignupRes struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
