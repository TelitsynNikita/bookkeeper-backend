package todo

type Request struct {
	Id          int      `json:"id"`
	Purpose     string   `json:"purpose"`
	Description string   `json:"description"`
	Status      string   `json:"status"`
	Amount      float64  `json:"amount"`
	Images      []string `json:"images"`
	UserId      int      `json:"user_id"`
}
