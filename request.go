package todo

type Request struct {
	Id           int      `json:"id" db:"id"`
	Purpose      string   `json:"purpose" db:"purpose"`
	Description  string   `json:"description" db:"description"`
	Status       string   `json:"status" db:"status"`
	Amount       float64  `json:"amount" db:"amount"`
	Images       []string `json:"images" db:"images"`
	UserId       int      `json:"user_id" db:"user_id"`
	BookkeeperId int      `json:"bookkeeper_id" db:"bookkeeper_id"`
}

type AllRequests struct {
	Id           int     `json:"id" db:"id"`
	Purpose      string  `json:"purpose" db:"purpose"`
	Amount       float64 `json:"amount" db:"amount"`
	Status       string  `json:"status" db:"status"`
	FullName     string  `json:"full_name" db:"full_name"`
	BookkeeperId int     `json:"bookkeeper_id" db:"bookkeeper_id"`
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
	CreatedAt   string   `json:"created_at" db:"created_at"`
}

type UpdateRequestStatus struct {
	Status    string `json:"status" db:"status"`
	RequestId int    `json:"request_id" db:"id"`
}

type UpdateRequestBookkeeperId struct {
	BookkeeperId int `json:"bookkeeper_id" db:"id"`
	RequestId    int `json:"request_id" db:"id"`
}
