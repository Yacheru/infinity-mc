package entities

type UserResponse struct {
	User   User   `json:"user"`
	Tokens Tokens `json:"tokens"`
}

type User struct {
	UserID   string `json:"user_id" db:"uuid" swaggerignore:"true"`
	Nickname string `json:"nickname" binding:"required" db:"nickname"`
	Role     string `json:"role" db:"role"`
	Email    string `json:"email" binding:"required,email" db:"email"`
	Password string `json:"password" binding:"required" db:"password"`
	IpAddr   string `json:"ip_addr" db:"ip" swaggerignore:"true"`
}

type UserLogin struct {
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
	IpAddress string `json:"ip_address" swaggerignore:"true"`
}

type SendCode struct {
	Email string `json:"email" binding:"required,email"`
}
