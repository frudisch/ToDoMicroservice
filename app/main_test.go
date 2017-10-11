package app_test

import (
	"os"
	"testing"

	"GoBlogEntry/app"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
)

var a app.App

func TestMain(m *testing.M) {
	os.Setenv("TEST_DB_CONNECTION", "localhost:5432")
	os.Setenv("TEST_DB_USERNAME", "go_user")
	os.Setenv("TEST_DB_PASSWORD", "go_user_passwd")
	os.Setenv("TEST_DB_NAME", "todo_test")

	a = app.App{}
	a.Initialize(
		os.Getenv("TEST_DB_CONNECTION"),
		os.Getenv("TEST_DB_USERNAME"),
		os.Getenv("TEST_DB_PASSWORD"),
		os.Getenv("TEST_DB_NAME"))

	ensureTableExists()

	code := m.Run()

	clearTable()

	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM todo")
	a.DB.Exec("ALTER SEQUENCE todo_id_seq RESTART WITH 1")
}

func TestEmptyTable(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/todos", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	body := response.Body.String()

	require.Equal(t, body, "[]", "Expected an empty array. Got %s", body)
}

func TestGetNonExistentTodo(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/todo/11", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)

	require.Equal(t, m["error"], "Todo not found", "Expected the 'error' key of the response to be set to 'Todo not found'. Got '%s'", m["error"])
}

func TestCreateTodo(t *testing.T) {
	clearTable()

	payload := []byte(`{"name": "GO REST API", "description":"Setup Go Rest API for Blog entry", "dueTo": 1507273200000}`)

	req, _ := http.NewRequest("POST", "/todo", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	require.Equal(t, m["name"], "GO REST API", "Expected todo name to be 'GO REST API'. Got '%v'", m["name"])

	require.Equal(t, m["description"], "Setup Go Rest API for Blog entry", "Expected todo description to be 'Setup Go Rest API for Blog entry'. Got '%v'", m["description"])

	// the id is compared to 1.0 because JSON unmarshaling converts numbers to
	// floats, when the target is a map[string]interface{}
	require.Equal(t, m["id"], 1.0, "Expected Todo ID to be '1'. Got '%v'", m["id"])
}

func TestGetTodo(t *testing.T) {
	clearTable()
	addTodo(1)

	req, _ := http.NewRequest("GET", "/todo/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestUpdateTodo(t *testing.T) {
	clearTable()
	addTodo(1)

	req, _ := http.NewRequest("GET", "/todo/1", nil)
	response := executeRequest(req)
	var originalTodo map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &originalTodo)

	payload := []byte(`{"name": "GO REST API", "description":"Setup Go Rest API for Blog entry", "dueTo": 1507273200000}`)

	req, _ = http.NewRequest("PUT", "/todo/1", bytes.NewBuffer(payload))
	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	require.Equal(t, m["id"], originalTodo["id"], "Expected the id to remain the same (%v). Got %v", originalTodo["id"], m["id"])

	require.NotEqual(t, m["name"], originalTodo["name"], "Expected the name to change from '%v' to '%v'. Got '%v'", originalTodo["name"], m["name"], m["name"])

	require.NotEqual(t, m["description"], originalTodo["description"], "Expected the price to change from '%v' to '%v'. Got '%v'", originalTodo["description"], m["description"], m["description"])
}

func addTodo(count int) {
	if count < 1 {
		count = 1
	}

	for i := 0; i < count; i++ {
		a.DB.Exec("INSERT INTO todo(NAME, DESCRIPTION, DUE_TO) VALUES($1, $2, $3)", "Todo name "+strconv.Itoa(i),
			"Todo description "+strconv.Itoa(i), (i+1.0)*1000000)
	}
}

func checkResponseCode(t *testing.T, expected, actual int) {
	require.Equal(t, expected, actual, "Expected response code %d. Got %d\n", expected, actual)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS todo
(
id SERIAL,
name TEXT NOT NULL,
description TEXT NOT NULL,
dueTo BIGINT NOT NULL,
CONSTRAINT todos_pkey PRIMARY KEY (id)
)`
