package routes

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"

	"github.com/ShayeGun/go-server/internal/storage/memory"
	"github.com/ShayeGun/go-server/models"
)

func TestGetUserHandler(t *testing.T) {
	// Create an instance of the application
	store := memory.NewRepository()

	ext := &ExternalDependencies{
		store,
	}

	uc := NewUserRoutes(ext)

	user := models.User{
		ID:       "2ac7231d-ba1b-42d9-8410-6e3fb74a3fbb",
		Email:    "test-mail",
		Password: "test-pass",
	}

	ext.GetUserTable().Add(user)

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodGet, "/2ac7231d-ba1b-42d9-8410-6e3fb74a3fbb", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("uid", "2ac7231d-ba1b-42d9-8410-6e3fb74a3fbb")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Call the handler function
	handler := http.HandlerFunc(uc.getUser())
	handler.ServeHTTP(rr, req)

	resUser := &models.User{}
	json.NewDecoder(rr.Body).Decode(resUser)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := user
	if *resUser != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestAddUserHandler(t *testing.T) {
	// Create an instance of the application
	store := memory.NewRepository()

	ext := &ExternalDependencies{
		store,
	}

	uc := NewUserRoutes(ext)

	data := map[string]string{
		"id":       "2ac7231d-ba1b-42d9-8410-6e3fb74a3fbb",
		"email":    "test-email",
		"password": "test-password",
	}

	jsonBody, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("Could not marshal request body: %v", err)
	}

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodGet, "", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Call the handler function
	handler := http.HandlerFunc(uc.addUser())
	handler.ServeHTTP(rr, req)

	resUser := &models.User{}
	json.NewDecoder(rr.Body).Decode(resUser)

	// Check the status code
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Check the response body
	expected := models.User{
		ID:       "2ac7231d-ba1b-42d9-8410-6e3fb74a3fbb",
		Email:    "test-email",
		Password: "test-password",
	}
	if *resUser != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestUpdateUserHandler(t *testing.T) {
	// Create an instance of the application
	store := memory.NewRepository()

	ext := &ExternalDependencies{
		store,
	}

	uc := NewUserRoutes(ext)

	user := models.User{
		ID:       "2ac7231d-ba1b-42d9-8410-6e3fb74a3fbb",
		Email:    "test-email",
		Password: "test-password",
	}

	data := map[string]string{
		"email":    "test-changed-email",
		"password": "test-changed-password",
	}

	ext.GetUserTable().Add(user)

	jsonBody, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("Could not marshal request body: %v", err)
	}

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodGet, "/2ac7231d-ba1b-42d9-8410-6e3fb74a3fbb", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("uid", "2ac7231d-ba1b-42d9-8410-6e3fb74a3fbb")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Call the handler function
	handler := http.HandlerFunc(uc.updateUser())
	handler.ServeHTTP(rr, req)

	resUser := &models.User{}
	json.NewDecoder(rr.Body).Decode(resUser)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}

	// Check the response body
	expected := models.User{
		ID:       "2ac7231d-ba1b-42d9-8410-6e3fb74a3fbb",
		Email:    "test-email",
		Password: "test-changed-password",
	}
	if *resUser != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", *resUser, expected)
	}
}

func TestDeleteUserHandler(t *testing.T) {
	// Create an instance of the application
	store := memory.NewRepository()

	ext := &ExternalDependencies{
		store,
	}

	uc := NewUserRoutes(ext)

	user := models.User{
		ID:       "2ac7231d-ba1b-42d9-8410-6e3fb74a3fbb",
		Email:    "test-mail",
		Password: "test-pass",
	}

	repo := ext.GetUserTable()
	repo.Add(user)

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodGet, "/2ac7231d-ba1b-42d9-8410-6e3fb74a3fbb", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("uid", "2ac7231d-ba1b-42d9-8410-6e3fb74a3fbb")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Call the handler function
	handler := http.HandlerFunc(uc.deleteUser())
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if _, err := repo.GetById(user.GetID()); err == nil {
		t.Errorf("user expected to get deleted but got found")
	}
}
