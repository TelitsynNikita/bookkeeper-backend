package todo

type User struct {
	Id       int    `json:"-" db:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type SignInUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpUser struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
