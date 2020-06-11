package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"user-management-service/api"
	"user-management-service/config"
	"user-management-service/database"
	"user-management-service/store"

	"github.com/gorilla/mux"
)

var router *mux.Router

func TestMain(m *testing.M) {
	SetupTestServer()
	os.Exit(m.Run())
}

func SetupTestServer() {
	config := config.Config{
		HttpPort: "8080",
		Database: config.Database{
			DbHost: "localhost",
			DbPort: "48000",
			DbUser: "user-management-service",
			DbPass: "user-management-service",
			DbName: "user-management-service",
		},
	}

	db := database.NewPostgres(config.Database.DbHost, config.Database.DbPort,
		config.Database.DbUser, config.Database.DbPass, config.Database.DbName)

	store := store.NewStore(db)

	server := api.NewServer(store)

	router = api.NewRouter(server)

	return
}

func TestCreateGroup(t *testing.T) {
	var jsonStr = []byte(`{"name": "programmers"}`)
	req, err := http.NewRequest(http.MethodPost, "/groups", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != 200 {
		t.Errorf("Error expecting 200, got %d", rr.Code)
	}
}

func TestCreateGroupBadJSON(t *testing.T) {
	var jsonStr = []byte(``)
	req, err := http.NewRequest(http.MethodPost, "/groups", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != 400 {
		t.Errorf("Error expecting 400, got %d", rr.Code)
	}
}

func TestCreateGroupNoName(t *testing.T) {
	var jsonStr = []byte(`{}`)
	req, err := http.NewRequest(http.MethodPost, "/groups", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != 400 {
		t.Errorf("Error expecting 400, got %d", rr.Code)
	}
}

func TestGetGroups(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/groups", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != 200 {
		t.Errorf("Error expecting 200, got %d", rr.Code)
	}
}

func TestGetGroup(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/groups/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != 200 {
		t.Errorf("Error expecting 200, got %d", rr.Code)
	}
}

func TestGetGroupBadInput(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/groups/doctors", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != 400 {
		t.Errorf("Error expecting 400, got %d", rr.Code)
	}
}

func TestGetGroupNoEntries(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/groups/100000", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != 404 {
		t.Errorf("Error expecting 404, got %d", rr.Code)
	}
}

func TestModifyGroup(t *testing.T) {
	var jsonStr = []byte(`{"name": "doctors"}`)
	req, err := http.NewRequest(http.MethodPut, "/groups/1", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != 200 {
		t.Errorf("Error expecting 200, got %d", rr.Code)
	}
}

func TestModifyGroupBadJSON(t *testing.T) {
	var jsonStr = []byte(``)
	req, err := http.NewRequest(http.MethodPut, "/groups/1", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != 400 {
		t.Errorf("Error expecting 400, got %d", rr.Code)
	}
}

func TestModifyGroupNoName(t *testing.T) {
	var jsonStr = []byte(`{}`)
	req, err := http.NewRequest(http.MethodPut, "/groups/1", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != 400 {
		t.Errorf("Error expecting 400, got %d", rr.Code)
	}
}

func TestCreateUser(t *testing.T) {
	var jsonStr = []byte(`{"group_id": 1, "email": "jack@gmail.com", "password": "jack5", "name": "Jack"}`)
	req, err := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != 200 {
		t.Errorf("Error expecting 200, got %d", rr.Code)
	}
}

func TestCreateUserBadJSON(t *testing.T) {
	var jsonStr = []byte(``)
	req, err := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != 400 {
		t.Errorf("Error expecting 400, got %d", rr.Code)
	}
}

func TestCreateUserNoGroupID(t *testing.T) {
	var jsonStr = []byte(`{"email": "jack@gmail.com", "password": "jack5", "name": "Jack"}`)
	req, err := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != 400 {
		t.Errorf("Error expecting 400, got %d", rr.Code)
	}
}

func TestCreateUserNoEmail(t *testing.T) {
	var jsonStr = []byte(`{"group_id": 1, "password": "jack5", "name": "Jack"}`)
	req, err := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != 400 {
		t.Errorf("Error expecting 400, got %d", rr.Code)
	}
}

func TestCreateUserNoPassword(t *testing.T) {
	var jsonStr = []byte(`{"group_id": 1, "email": "jack@gmail.com", "name": "Jack"}`)
	req, err := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != 400 {
		t.Errorf("Error expecting 400, got %d", rr.Code)
	}
}

func TestCreateUserNoName(t *testing.T) {
	var jsonStr = []byte(`{"group_id": 1, "email": "jack@gmail.com", "password": "jack5"}`)
	req, err := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != 400 {
		t.Errorf("Error expecting 400, got %d", rr.Code)
	}
}

func TestGetUsers(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != 200 {
		t.Errorf("Error expecting 200, got %d", rr.Code)
	}
}

func TestGetUser(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != 200 {
		t.Errorf("Error expecting 200, got %d", rr.Code)
	}
}

func TestGetUserBadInput(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/users/joe", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != 400 {
		t.Errorf("Error expecting 400, got %d", rr.Code)
	}
}

func TestGetUserNoEntries(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/users/100000", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != 404 {
		t.Errorf("Error expecting 404, got %d", rr.Code)
	}
}

func TestModifyUser(t *testing.T) {
	var jsonStr = []byte(`{"group_id": 1, "email": "joe@gmail.com", "password": "joe5", "name": "Joe"}`)
	req, err := http.NewRequest(http.MethodPut, "/users/1", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != 200 {
		t.Errorf("Error expecting 200, got %d", rr.Code)
	}
}

func TestModifyUserBadJSON(t *testing.T) {
	var jsonStr = []byte(``)
	req, err := http.NewRequest(http.MethodPut, "/users/1", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != 400 {
		t.Errorf("Error expecting 400, got %d", rr.Code)
	}
}

func TestModifyUserNoGroupID(t *testing.T) {
	var jsonStr = []byte(`{"email": "joe@gmail.com", "password": "joe5", "name": "Joe"}`)
	req, err := http.NewRequest(http.MethodPut, "/users/1", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != 400 {
		t.Errorf("Error expecting 400, got %d", rr.Code)
	}
}

func TestModifyUserNoEmail(t *testing.T) {
	var jsonStr = []byte(`{"group_id": 1, "password": "joe5", "name": "Joe"}`)
	req, err := http.NewRequest(http.MethodPut, "/users/1", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != 400 {
		t.Errorf("Error expecting 400, got %d", rr.Code)
	}
}

func TestModifyUserNoPassword(t *testing.T) {
	var jsonStr = []byte(`{"group_id": 1, "email": "joe@gmail.com", "name": "Joe"}`)
	req, err := http.NewRequest(http.MethodPut, "/users/1", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != 400 {
		t.Errorf("Error expecting 400, got %d", rr.Code)
	}
}

func TestModifyUserNoName(t *testing.T) {
	var jsonStr = []byte(`{"group_id": 1, "email": "joe@gmail.com", "password": "jo5"}`)
	req, err := http.NewRequest(http.MethodPut, "/users/1", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != 400 {
		t.Errorf("Error expecting 400, got %d", rr.Code)
	}
}

func TestGetGroupWithUsers(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/groups/1/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != 200 {
		t.Errorf("Error expecting 200, got %d", rr.Code)
	}
}
