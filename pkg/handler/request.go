package handler

import (
	"encoding/json"
	todo "github.com/TelitsynNikita/bookkeeper-backend"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (h *Handler) GetAllRequests(w http.ResponseWriter, r *http.Request) {
	requests, err := h.services.RequestList.GetAll()
	if err != nil {
		NewErrorResponse(w, err.Error())
		return
	}

	b, _ := json.Marshal(requests)
	w.Write(b)
}

func (h *Handler) GetOneRequest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	requestId, _ := strconv.Atoi(vars["id"])

	request, err := h.services.RequestList.GetOne(requestId)
	if err != nil {
		NewErrorResponse(w, "Такой запрос не существует")
		return
	}

	b, _ := json.Marshal(request)
	w.Write(b)
}

func (h *Handler) CreateRequest(w http.ResponseWriter, r *http.Request) {
	userId, err := h.services.ParseToken(r.Header.Get(authorizationHeader))
	if err != nil {
		NewErrorResponse(w, err.Error())
		return
	}

	if err := r.ParseForm(); err != nil {
		NewErrorResponse(w, err.Error())
		return
	}

	var input todo.Request

	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		NewErrorResponse(w, err.Error())
		return
	}

	id, err := h.services.RequestList.Create(userId, input)
	if err != nil {
		NewErrorResponse(w, "Некорректные данные")
		return
	}

	res, err := json.Marshal(map[string]interface{}{
		"id": id,
	})
	if err != nil {
		NewErrorResponse(w, err.Error())
		return
	}

	w.Write(res)

}

func (h *Handler) UpdateRequest(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		NewErrorResponse(w, err.Error())
		return
	}

	var input []todo.UpdateRequestStatus

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		NewErrorResponse(w, err.Error())
		return
	}

	for _, request := range input {
		err = h.services.RequestList.Update(request)
		if err != nil {
			NewErrorResponse(w, "Невозможно изменить статус")
			return
		}
	}

	res, _ := json.Marshal(map[string]interface{}{
		"message": "Успешно обновлено",
	})

	w.Write(res)
}

func (h *Handler) DeleteRequest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	requestId, err := strconv.Atoi(vars["id"])
	if err != nil {
		NewErrorResponse(w, err.Error())
		return
	}

	err = h.services.RequestList.DeleteOne(requestId)
	if err != nil {
		NewErrorResponse(w, "Невозможно удалить несуществующую заявку")
		return
	}

	res, _ := json.Marshal(map[string]interface{}{
		"message": "Успешно удалено",
	})

	w.Write(res)
}

func (h *Handler) UpdateRequestBookkeeperId(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		NewErrorResponse(w, err.Error())
		return
	}

	var input todo.UpdateRequestBookkeeperId

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		NewErrorResponse(w, err.Error())
		return
	}

	err = h.services.UpdateRequestBookkeeperId(input)
	if err != nil {
		NewErrorResponse(w, "Невозможно обновить")
		return
	}

	res, _ := json.Marshal(map[string]interface{}{
		"message": "Успешно обновлено",
	})

	w.Write(res)
}
