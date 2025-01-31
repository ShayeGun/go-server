package routes

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"

	"github.com/ShayeGun/go-server/internal/common"
	"github.com/ShayeGun/go-server/internal/service"
	"github.com/ShayeGun/go-server/internal/storage/memory"
	"github.com/ShayeGun/go-server/models"
)

func TestGetUserHandler(t *testing.T) {
	// ARRANGE
	dep := common.ExternalDependencies{
		RepositoryInterface: memory.NewRepository(),
	}

	services, _ := service.NewService(dep)
	uc := NewUserRoutes(services.GetUserService())

	user := models.User{
		ID:       "2ac7231d-ba1b-42d9-8410-6e3fb74a3fbb",
		Email:    "test-mail",
		Password: "test-pass",
	}

	dep.GetUserTable().Add(user)

	req, err := http.NewRequest(http.MethodGet, "/2ac7231d-ba1b-42d9-8410-6e3fb74a3fbb", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("uid", "2ac7231d-ba1b-42d9-8410-6e3fb74a3fbb")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// ACT
	handler := http.HandlerFunc(uc.getUser())
	handler.ServeHTTP(rr, req)

	resUser := &models.User{}
	json.NewDecoder(rr.Body).Decode(resUser)

	// ASSERT
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := user
	if *resUser != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestAddUserHandler(t *testing.T) {
	// ARRANGE
	dep := common.ExternalDependencies{
		RepositoryInterface: memory.NewRepository(),
	}
	services, _ := service.NewService(dep)
	uc := NewUserRoutes(services.GetUserService())

	data := map[string]string{
		"id":       "2ac7231d-ba1b-42d9-8410-6e3fb74a3fbb",
		"email":    "test-email",
		"password": "test-password",
	}

	jsonBody, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("Could not marshal request body: %v", err)
	}

	req, err := http.NewRequest(http.MethodGet, "", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// ACT
	handler := http.HandlerFunc(uc.addUser())
	handler.ServeHTTP(rr, req)

	resUser := &models.User{}
	json.NewDecoder(rr.Body).Decode(resUser)

	// ASSERT
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

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
	// ARRANGE
	dep := common.ExternalDependencies{
		RepositoryInterface: memory.NewRepository(),
	}
	services, _ := service.NewService(dep)
	uc := NewUserRoutes(services.GetUserService())

	user := models.User{
		ID:       "2ac7231d-ba1b-42d9-8410-6e3fb74a3fbb",
		Email:    "test-email",
		Password: "test-password",
	}

	data := map[string]string{
		"email":    "test-changed-email",
		"password": "test-changed-password",
	}

	dep.GetUserTable().Add(user)

	jsonBody, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("Could not marshal request body: %v", err)
	}

	req, err := http.NewRequest(http.MethodGet, "/2ac7231d-ba1b-42d9-8410-6e3fb74a3fbb", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("uid", "2ac7231d-ba1b-42d9-8410-6e3fb74a3fbb")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// ACT
	handler := http.HandlerFunc(uc.updateUser())
	handler.ServeHTTP(rr, req)

	resUser := &models.User{}
	json.NewDecoder(rr.Body).Decode(resUser)

	// ASSERT
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}

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
	// ARRANGE
	dep := common.ExternalDependencies{
		RepositoryInterface: memory.NewRepository(),
	}
	services, _ := service.NewService(dep)
	uc := NewUserRoutes(services.GetUserService())

	user := models.User{
		ID:       "2ac7231d-ba1b-42d9-8410-6e3fb74a3fbb",
		Email:    "test-mail",
		Password: "test-pass",
	}

	repo := dep.GetUserTable()
	repo.Add(user)

	req, err := http.NewRequest(http.MethodGet, "/2ac7231d-ba1b-42d9-8410-6e3fb74a3fbb", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("uid", "2ac7231d-ba1b-42d9-8410-6e3fb74a3fbb")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// ACT
	handler := http.HandlerFunc(uc.deleteUser())
	handler.ServeHTTP(rr, req)

	// ASSERT
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if _, err := repo.GetById(user.GetID()); err == nil {
		t.Errorf("user expected to get deleted but got found")
	}
}
