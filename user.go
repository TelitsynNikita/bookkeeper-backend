package todo

type User struct {
	Id       int    `json:"-"`
	Fullname string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
