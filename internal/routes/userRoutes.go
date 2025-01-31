package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ShayeGun/go-server/internal/service"
	"github.com/ShayeGun/go-server/internal/storage/memory"
	"github.com/ShayeGun/go-server/internal/util"
	"github.com/ShayeGun/go-server/models"
	"github.com/go-chi/chi/v5"
)

type RepositoryInterface interface {
	GetUserTable() service.UserRepositoryInterface
}

type ExternalDependencies struct {
	RepositoryInterface
	//logger
	// cache
}

type UserRoutes struct {
	dep *ExternalDependencies
}

func NewUserRoutes(dep *ExternalDependencies) UserRoutes {
	return UserRoutes{
		dep: dep,
	}
}

func (u *UserRoutes) SetupUserRoutes(r *chi.Mux) *chi.Mux {
	r.Get("/v1/users/{uid}", u.getUser())
	r.Post("/v1/users", u.addUser())
	r.Delete("/v1/users/{uid}", u.deleteUser())
	r.Patch("/v1/users/{uid}", u.updateUser())

	return r
}

func (u *UserRoutes) getUser() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		us, err := service.NewUserService(
			service.WithUserRepository(u.dep.GetUserTable()),
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

func (u *UserRoutes) addUser() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		us, err := service.NewUserService(
			service.WithUserRepository(u.dep.GetUserTable()),
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

func (u *UserRoutes) updateUser() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		us, err := service.NewUserService(
			service.WithUserRepository(u.dep.GetUserTable()),
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

func (u *UserRoutes) deleteUser() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		us, err := service.NewUserService(
			service.WithUserRepository(u.dep.GetUserTable()),
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
