package handler

import (
	"github.com/TelitsynNikita/bookkeeper-backend/pkg/service"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() http.Handler {
	router := mux.NewRouter()

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"OPTIONS", "GET", "POST", "PUT", "DELETE", "PATCH"})
	origins := handlers.AllowedOrigins([]string{"*"})

	//router.Use(h.userIdentity)

	router.HandleFunc("/auth/sign-up", h.signUp).Methods("POST")
	router.HandleFunc("/auth/sign-in", h.signIn).Methods("POST")
	router.HandleFunc("/api/requests", h.GetAllRequests).Methods("GET")
	router.HandleFunc("/api/requests/{id}", h.GetOneRequest).Methods("GET")
	router.HandleFunc("/api/requests", h.CreateRequest).Methods("POST")
	router.HandleFunc("/api/requests/{id}", h.DeleteRequest).Methods("DELETE")
	router.HandleFunc("/api/requests", h.UpdateRequest).Methods("PATCH")

	return handlers.CORS(headers, methods, origins)(router)
}
