package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(srv *Server) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/users", srv.GetUsers).Methods(http.MethodGet)
	r.HandleFunc("/users/{userID}", srv.GetUser).Methods(http.MethodGet)
	r.HandleFunc("/users", srv.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/users/{userID}", srv.ModifyUser).Methods(http.MethodPut)
	r.HandleFunc("/users/{userID}", srv.RemoveUser).Methods(http.MethodDelete)

	r.HandleFunc("/groups", srv.GetGroups).Methods(http.MethodGet)
	r.HandleFunc("/groups/{groupID}", srv.GetGroup).Methods(http.MethodGet)
	r.HandleFunc("/groups/{groupID}/users", srv.GetGroupWithUsers).Methods(http.MethodGet)
	r.HandleFunc("/groups", srv.CreateGroup).Methods(http.MethodPost)
	r.HandleFunc("/groups/{groupID}", srv.ModifyGroup).Methods(http.MethodPut)
	r.HandleFunc("/groups/{groupID}", srv.RemoveGroup).Methods(http.MethodDelete)

	return r
}
