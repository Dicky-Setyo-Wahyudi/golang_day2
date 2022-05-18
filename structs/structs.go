package structs

// Users is a representation of a post
type Users struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	NoTelepon string `json:"notelepon"`
	Role      string `json:"role"`
	Status    bool   `json:"status"`
}

type UsersLogin struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Product struct {
	ID         int    `json:"id"`
	NamaProduk string `json:"namaproduk"`
	Deskripsi  string `json:"deskripsi"`
	Status     bool   `json:"status"`
}

// Result is an array of post
type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
