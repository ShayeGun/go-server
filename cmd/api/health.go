package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/ShayeGun/go-server/internal/service"
	"github.com/ShayeGun/go-server/internal/storage/memory"
	"github.com/ShayeGun/go-server/internal/util"
	"github.com/ShayeGun/go-server/models"
	"github.com/go-chi/chi/v5"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func (app *application) slowHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second * 5)
	w.Write([]byte("slow done"))
}

func (app *application) getUser(e *externalDependencies) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		us, err := service.NewUserService(
			service.WithUserRepository(e.GetUserTable()),
		)

		if err != nil {
			log.Println(err)
			util.WriteJSONError(w, http.StatusInternalServerError, "internal service error")
			return
		}
		uidStr := chi.URLParam(r, "uid")

		u, err := us.UserRepository.GetById(uidStr)
		if err != nil {
			util.WriteJSONError(w, http.StatusNotFound, err.Error())
			return
		}

		util.WriteJSONResponse(w, http.StatusOK, u)
	}
}

func (app *application) addUser(e *externalDependencies) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		us, err := service.NewUserService(
			service.WithUserRepository(e.GetUserTable()),
		)

		if err != nil {
			log.Println(err)
			util.WriteJSONError(w, http.StatusInternalServerError, "internal service error")
			return
		}

		var user models.User

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			log.Println("invalid params")
			util.WriteJSONError(w, http.StatusBadRequest, "invalid params")
			return
		}

		u, err := us.UserRepository.Add(user)

		if err != nil {
			log.Println(memory.ErrUserAlreadyExists)
			util.WriteJSONError(w, http.StatusBadRequest, err.Error())
			return
		}

		util.WriteJSONResponse(w, http.StatusCreated, u)
	}
}

func (app *application) updateUser(e *externalDependencies) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		us, err := service.NewUserService(
			service.WithUserRepository(e.GetUserTable()),
		)

		if err != nil {
			log.Println(err)
			util.WriteJSONError(w, http.StatusInternalServerError, "internal service error")
			return
		}

		uidStr := chi.URLParam(r, "uid")

		user := models.User{
			ID: uidStr,
		}

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			log.Println("invalid params")
			util.WriteJSONError(w, http.StatusBadRequest, "invalid params")
			return
		}

		u, err := us.UserRepository.Update(user)

		if err != nil {
			log.Println(memory.ErrUserAlreadyExists)
			util.WriteJSONError(w, http.StatusBadRequest, err.Error())
			return
		}

		util.WriteJSONResponse(w, http.StatusOK, u)
	}
}

func (app *application) deleteUser(e *externalDependencies) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		us, err := service.NewUserService(
			service.WithUserRepository(e.GetUserTable()),
		)

		if err != nil {
			log.Println(err)
			util.WriteJSONError(w, http.StatusInternalServerError, "internal service error")
			return
		}
		uidStr := chi.URLParam(r, "uid")

		if err := us.UserRepository.Delete(uidStr); err != nil {
			util.WriteJSONError(w, http.StatusNotFound, err.Error())
			return
		}

		w.Write([]byte("User Deleted."))

	}
}
