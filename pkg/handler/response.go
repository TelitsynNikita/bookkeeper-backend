package handler

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ErrorMessage struct {
	Message string `json:"message"`
}

func NewErrorResponse(w http.ResponseWriter, message string) {
	b, _ := json.Marshal(map[string]interface{}{
		"message": message,
	})
	w.Write([]byte(b))
	logrus.Error(message)
}
