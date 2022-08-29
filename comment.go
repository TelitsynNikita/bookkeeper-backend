package todo

type Comment struct {
	Id           int    `json:"id"`
	Message      string `json:"message"`
	RequestId    int    `json:"request_id"`
	BookkeeperId int    `json:"bookkeeper_id"`
}
