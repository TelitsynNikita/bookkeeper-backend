package todo

type Request struct {
	Id          int      `json:"id" db:"id"`
	Purpose     string   `json:"purpose" db:"purpose"`
	Description string   `json:"description" db:"description"`
	Status      string   `json:"status" db:"status"`
	Amount      float64  `json:"amount" db:"amount"`
	Images      []string `json:"images" db:"images"`
	UserId      int      `json:"user_id" db:"user_id"`
}

type AllRequests struct {
	Id       int     `json:"id" db:"id"`
	Purpose  string  `json:"purpose" db:"purpose"`
	Amount   float64 `json:"amount" db:"amount"`
	Status   string  `json:"status" db:"status"`
	FullName string  `json:"full_name" db:"full_name"`
}

type OneRequest struct {
	Id          int      `json:"id" db:"id"`
	Purpose     string   `json:"purpose" db:"purpose"`
	Description string   `json:"description" db:"description"`
	Status      string   `json:"status" db:"status"`
	Amount      float64  `json:"amount" db:"amount"`
	Images      []string `json:"images" db:"images"`
	UserId      int      `json:"user_id" db:"user_id"`
	FullName    string   `json:"full_name" db:"full_name"`
}
