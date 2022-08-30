package repository

import (
	"fmt"
	todo "github.com/TelitsynNikita/bookkeeper-backend"
	"github.com/jmoiron/sqlx"
)

type RequestListPostgres struct {
	db *sqlx.DB
}

func NewRequestListPostgres(db *sqlx.DB) *RequestListPostgres {
	return &RequestListPostgres{db: db}
}

func (r *RequestListPostgres) Create(userId int, request todo.Request) (int, error) {
	var id int
	fmt.Println(request.Images)
	createRequestQuery := fmt.Sprintf("INSERT INTO %s (purpose, description, amount, user_id) values ($1, $2, $3, $4) RETURNING id", requestsTable)
	row := r.db.QueryRow(createRequestQuery, request.Purpose, request.Description, request.Amount, userId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *RequestListPostgres) GetAll(userId int) ([]todo.AllRequests, error) {
	var requests []todo.AllRequests
	query := fmt.Sprintf("SELECT tl.id, tl.purpose, tl.amount, tl.status, ul.full_name FROM %s tl JOIN %s ul on tl.user_id = ul.id WHERE ul.id = $1",
		requestsTable, usersTable)
	err := r.db.Select(&requests, query, userId)

	return requests, err
}

func (r *RequestListPostgres) GetOne(requestId int) (todo.OneRequest, error) {
	var request todo.OneRequest
	query := fmt.Sprintf("SELECT tl.id, tl.purpose, tl.amount, tl.status, ul.full_name FROM %s tl JOIN %s ul on tl.user_id = ul.id WHERE tl.id = $1",
		requestsTable, usersTable)
	err := r.db.Get(&request, query, requestId)

	return request, err
}