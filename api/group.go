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

func (srv *Server) GetGroups(w http.ResponseWriter, req *http.Request) {
	getGroups, err := srv.store.GetGroups()
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(HandlerError("Internal server error")))
		return
	}

	groups, err := json.Marshal(getGroups)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(HandlerError("Internal server error")))
		return
	}

	w.Write(groups)
}

func (srv *Server) GetGroup(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	groupID := vars["groupID"]
	intGroupID, err := strconv.Atoi(groupID)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(HandlerError("Bad input")))
		return
	}

	getGroup, err := srv.store.GetGroup(intGroupID)
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

	group, err := json.Marshal(getGroup)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(HandlerError("Internal server error")))
		return
	}

	w.Write(group)
}

func (srv *Server) GetGroupWithUsers(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	groupID := vars["groupID"]
	intGroupID, err := strconv.Atoi(groupID)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(HandlerError("Bad input")))
		return
	}

	result, err := srv.store.GetGroupWithUsers(intGroupID)
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

	out, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(HandlerError("Internal server error")))
		return
	}

	w.Write(out)
}

func (srv *Server) CreateGroup(w http.ResponseWriter, req *http.Request) {
	inputJSON, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
	}

	var group model.Group

	errInputJSON := json.Unmarshal([]byte(inputJSON), &group)
	w.Header().Set("Content-Type", "application/json")
	if errInputJSON != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(HandlerError("Bad json!")))
		return
	}

	if group.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(HandlerError("Bad request, name is required!")))
		return
	}

	group, err = srv.store.CreateGroup(group.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newGroup, err := json.Marshal(group)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(newGroup)
}

func (srv *Server) ModifyGroup(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	groupID := vars["groupID"]
	intGroupID, err := strconv.Atoi(groupID)
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

	var group model.Group

	errInputJSON := json.Unmarshal([]byte(inputJSON), &group)
	if errInputJSON != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(HandlerError("Bad json!")))
		return
	}

	group.GroupID = intGroupID

	if group.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(HandlerError("Bad request, name is required!")))
		return
	}

	err = srv.store.ModifyGroup(intGroupID, group)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(HandlerError("No entries found")))
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	modifiedGroup, err := json.Marshal(group)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(modifiedGroup)
}

func (srv *Server) RemoveGroup(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	groupID := vars["groupID"]
	intGroupID, err := strconv.Atoi(groupID)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(HandlerError("Bad input")))
		return
	}

	err = srv.store.RemoveGroup(intGroupID)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(HandlerError("No entries found")))
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(HandlerError(err.Error())))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
