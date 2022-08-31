package handler

import (
	"encoding/json"
	todo "github.com/TelitsynNikita/bookkeeper-backend"
	"net/http"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		NewErrorResponse(w, err.Error())
		return
	}

	var user todo.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		NewErrorResponse(w, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(user)
	if err != nil {
		NewErrorResponse(w, err.Error())
		return
	}

	res, err := json.Marshal(map[string]interface{}{
		"id": id,
	})
	if err != nil {
		NewErrorResponse(w, err.Error())
		return
	}

	w.Write([]byte(res))
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		return
	}

	var user todo.SignInUser

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		NewErrorResponse(w, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(user.Email, user.Password)
	if err != nil {
		NewErrorResponse(w, err.Error())
		return
	}

	res, err := json.Marshal(map[string]interface{}{
		"token": token,
	})
	if err != nil {
		NewErrorResponse(w, err.Error())
		return
	}
	w.Write([]byte(res))
}
