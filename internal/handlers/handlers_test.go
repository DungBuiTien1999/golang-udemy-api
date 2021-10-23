package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

var tasksTest = []struct {
	testCase     string
	url          string
	expectedCode int
}{
	{
		testCase:     "valid case",
		url:          "/api/tasks/1",
		expectedCode: http.StatusOK,
	},
	{
		testCase:     "id is not an integer",
		url:          "/api/tasks/invalid",
		expectedCode: http.StatusBadRequest,
	},
	{
		testCase:     "error in connect to db",
		url:          "/api/tasks/2",
		expectedCode: http.StatusBadRequest,
	},
}

func TestRepository_TasksOfUser(t *testing.T) {
	for _, e := range tasksTest {
		req, _ := http.NewRequest("GET", e.url, nil)
		req.RequestURI = e.url

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(Repo.TasksOfUser)
		handler.ServeHTTP(rr, req)

		if rr.Code != e.expectedCode {
			t.Errorf("fail %s: got %d, expected %d", e.testCase, rr.Code, e.expectedCode)
		}
	}
}

var registerTest = []struct {
	testCase     string
	jsonData     []byte
	expectedCode int
}{
	{
		testCase: "valid case",
		jsonData: []byte(`
		{
			"username": "dungbui",
			"password": "samurai"
		}
	`),
		expectedCode: http.StatusCreated,
	},
	{
		testCase: "invalid input data",
		jsonData: []byte(`
		{
			"username": "dungbui"
		}
	`),
		expectedCode: http.StatusUnprocessableEntity,
	},
	{
		testCase: "error in connect to db",
		jsonData: []byte(`
		{
			"username": "ronando",
			"password": "dungbui"
		}
	`),
		expectedCode: http.StatusBadRequest,
	},
}

func TestRepository_Register(t *testing.T) {
	for _, e := range registerTest {
		jsonData := e.jsonData
		req, _ := http.NewRequest("POST", "/api/users/", bytes.NewBuffer(jsonData))

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(Repo.Register)
		handler.ServeHTTP(rr, req)

		if rr.Code != e.expectedCode {
			t.Errorf("fail %s: got %d, expected %d", e.testCase, rr.Code, e.expectedCode)
		}
	}
}

var addTaskTest = []struct {
	testCase     string
	jsonData     []byte
	expectedCode int
}{
	{
		testCase: "valid case",
		jsonData: []byte(`
		{
			"title": "learning golang",
			"user_id": 1
		}
	`),
		expectedCode: http.StatusCreated,
	},
	{
		testCase: "invalid input data",
		jsonData: []byte(`
		{
			"title": "learning golang"
		}
	`),
		expectedCode: http.StatusUnprocessableEntity,
	},
	{
		testCase: "error in connect to db",
		jsonData: []byte(`
		{
			"title": "learning golang",
			"user_id": 2
		}
	`),
		expectedCode: http.StatusBadRequest,
	},
}

func TestRepository_AddTask(t *testing.T) {
	for _, e := range addTaskTest {
		jsonData := e.jsonData
		req, _ := http.NewRequest("POST", "/api/tasks/", bytes.NewBuffer(jsonData))

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(Repo.AddTask)
		handler.ServeHTTP(rr, req)

		if rr.Code != e.expectedCode {
			t.Errorf("fail %s: got %d, expected %d", e.testCase, rr.Code, e.expectedCode)
		}
	}
}

var updateTaskTest = []struct {
	testCase     string
	url          string
	jsonData     []byte
	expectedCode int
}{
	{
		testCase: "valid case",
		url:      "/api/tasks/1",
		jsonData: []byte(`
		{
			"title": "learning golang",
			"user_id": 1,
			"complete": 1
		}
	`),
		expectedCode: http.StatusOK,
	},
	{
		testCase: "invalid id at url",
		url:      "/api/tasks/invalid",
		jsonData: []byte(`
		{
			"title": "learning golang",
			"user_id": 1,
			"complete": 1
		}
	`),
		expectedCode: http.StatusBadRequest,
	},
	{
		testCase: "invalid input data",
		url:      "/api/tasks/1",
		jsonData: []byte(`
		{
			"title": "learning golang"
		}
	`),
		expectedCode: http.StatusUnprocessableEntity,
	},
	{
		testCase: "error in connect to db",
		url:      "/api/tasks/2",
		jsonData: []byte(`
		{
			"title": "learning golang",
			"user_id": 1,
			"complete": 1
		}
	`),
		expectedCode: http.StatusBadRequest,
	},
}

func TestRepository_UpdateTask(t *testing.T) {
	for _, e := range updateTaskTest {
		jsonData := e.jsonData
		req, _ := http.NewRequest("PUT", e.url, bytes.NewBuffer(jsonData))
		req.RequestURI = e.url

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(Repo.UpdateTask)
		handler.ServeHTTP(rr, req)

		if rr.Code != e.expectedCode {
			t.Errorf("fail %s: got %d, expected %d", e.testCase, rr.Code, e.expectedCode)
		}
	}
}

var authenticationTest = []struct {
	testCase     string
	jsonData     []byte
	expectedCode int
}{
	{
		testCase: "valid case",
		jsonData: []byte(`
		{
			"username": "tiendung",
			"password": "tiendung"
		}
	`),
		expectedCode: http.StatusOK,
	},
	{
		testCase: "invalid input data",
		jsonData: []byte(`
		{
			"username": "dungbui"
		}
	`),
		expectedCode: http.StatusUnprocessableEntity,
	},
	{
		testCase: "username is non-existed",
		jsonData: []byte(`
		{
			"username": "ronando",
			"password": "dungbui"
		}
	`),
		expectedCode: http.StatusNotFound,
	},
	{
		testCase: "password wrong",
		jsonData: []byte(`
		{
			"username": "tiendung",
			"password": "dungbui1"
		}
	`),
		expectedCode: http.StatusBadRequest,
	},
	{
		testCase: "error in case update rfToken",
		jsonData: []byte(`
		{
			"username": "samurai",
			"password": "samurai"
		}
	`),
		expectedCode: http.StatusBadRequest,
	},
}

func TestRepository_Authentication(t *testing.T) {
	for _, e := range authenticationTest {
		jsonData := e.jsonData
		req, _ := http.NewRequest("POST", "/api/tasks/", bytes.NewBuffer(jsonData))

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(Repo.Authentication)
		handler.ServeHTTP(rr, req)

		if rr.Code != e.expectedCode {
			t.Errorf("fail %s: got %d, expected %d", e.testCase, rr.Code, e.expectedCode)
		}
	}
}
