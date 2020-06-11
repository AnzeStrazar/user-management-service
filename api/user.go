package api

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"user-management-service/model"

	"github.com/gorilla/mux"
)

func (srv *Server) GetUsers(w http.ResponseWriter, req *http.Request) {
	getUsers, err := srv.store.GetUsers()
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(HandlerError("Internal server error")))
		return
	}

	users, err := json.Marshal(getUsers)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(HandlerError("Internal server error")))
		return
	}

	w.Write(users)
}

func (srv *Server) GetUser(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	userID := vars["userID"]
	intUserID, err := strconv.Atoi(userID)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(HandlerError("Bad input")))
		return
	}

	getUser, err := srv.store.GetUser(intUserID)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(HandlerError("No entries found")))
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(HandlerError("Internal server error")))
		return
	}

	user, err := json.Marshal(getUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(HandlerError("Internal server error")))
		return
	}

	w.Write(user)
}

func (srv *Server) CreateUser(w http.ResponseWriter, req *http.Request) {
	inputJSON, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
	}

	var user model.User

	errInputJSON := json.Unmarshal([]byte(inputJSON), &user)
	w.Header().Set("Content-Type", "application/json")
	if errInputJSON != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(HandlerError("Bad json!")))
		return
	}

	if user.GroupID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(HandlerError("Bad request, group_id is required!")))
		return
	} else if user.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(HandlerError("Bad request, email is required!")))
		return
	} else if user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(HandlerError("Bad request, password is required!")))
		return
	} else if user.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(HandlerError("Bad request, name is required!")))
		return
	}

	user, err = srv.store.CreateUser(user.GroupID, user.Email, user.Password, user.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newUser, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(newUser)
}

func (srv *Server) ModifyUser(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	userID := vars["userID"]
	intUserID, err := strconv.Atoi(userID)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(HandlerError("Bad input")))
		return
	}

	inputJSON, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
	}

	var user model.User

	errInputJSON := json.Unmarshal([]byte(inputJSON), &user)
	if errInputJSON != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(HandlerError("Bad json!")))
		return
	}

	user.UserID = intUserID

	if user.GroupID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(HandlerError("Bad request, group_id is required!")))
		return
	} else if user.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(HandlerError("Bad request, email is required!")))
		return
	} else if user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(HandlerError("Bad request, password is required!")))
		return
	} else if user.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(HandlerError("Bad request, name is required!")))
		return
	}

	err = srv.store.ModifyUser(intUserID, user)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(HandlerError("No entries found")))
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

	modifiedUser, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(modifiedUser)
}

func (srv *Server) RemoveUser(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	userID := vars["userID"]
	intUserID, err := strconv.Atoi(userID)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(HandlerError("Bad input")))
		return
	}

	err = srv.store.RemoveUser(intUserID)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(HandlerError("No entries found")))
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(HandlerError("Internal server error")))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
