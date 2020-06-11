package api

import (
	"encoding/json"
	"user-management-service/store"
)

type Server struct {
	store *store.Store
}

func NewServer(store *store.Store) *Server {
	server := &Server{
		store: store,
	}

	return server
}

func HandlerError(errorText string) []byte {
	rsp := make(map[string]string)
	rsp["error"] = errorText
	jsonRsp, _ := json.Marshal(rsp)

	return jsonRsp
}
