package core

type User struct {
	Id          int    `json:"id" db:"id"`
	Email       string `json:"email" db:"email"`
	Password    string `json:"password" db:"password"`
	Username    string `json:"username" db:"username"`
	Description string `json:"description" db:"description"`
	Photo       string `json:"photo" db:"photo"`
	RoleId      int    `json:"role_id" db:"role_id"`
	IsBanned    bool   `json:"is_banned" db:"is_banned"`
	CreatedAt   string `json:"created_at" db:"created_at"`
}
