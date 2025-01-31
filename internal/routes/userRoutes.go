package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ShayeGun/go-server/internal/common"
	"github.com/ShayeGun/go-server/internal/storage/memory"
	"github.com/ShayeGun/go-server/internal/util"
	"github.com/ShayeGun/go-server/models"
	"github.com/go-chi/chi/v5"
)

type UserRoutes struct {
	userService common.UserServiceInterface
}

func NewUserRoutes(us common.UserServiceInterface) UserRoutes {
	return UserRoutes{
		userService: us,
	}
}

func (u *UserRoutes) SetupUserRoutes(r *chi.Mux) *chi.Mux {
	r.Get("/v1/users/{uid}", u.getUser())
	r.Post("/v1/users", u.addUser())
	r.Delete("/v1/users/{uid}", u.deleteUser())
	r.Patch("/v1/users/{uid}", u.updateUser())

	return r
}

func (ur *UserRoutes) getUser() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		uidStr := chi.URLParam(r, "uid")
		u, err := ur.userService.GetUser(uidStr)
		if err != nil {
			util.WriteJSONError(w, http.StatusNotFound, err.Error())
			return
		}

		util.WriteJSONResponse(w, http.StatusOK, u)
	}
}

func (ur *UserRoutes) addUser() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			log.Println("invalid params")
			util.WriteJSONError(w, http.StatusBadRequest, "invalid params")
			return
		}

		u, err := ur.userService.AddUser(user)

		if err != nil {
			log.Println(memory.ErrUserAlreadyExists)
			util.WriteJSONError(w, http.StatusBadRequest, err.Error())
			return
		}

		util.WriteJSONResponse(w, http.StatusCreated, u)
	}
}

func (ur *UserRoutes) updateUser() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		uidStr := chi.URLParam(r, "uid")

		user := models.User{
			ID: uidStr,
		}

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			log.Println("invalid params")
			util.WriteJSONError(w, http.StatusBadRequest, "invalid params")
			return
		}

		u, err := ur.userService.UpdateUser(user)

		if err != nil {
			log.Println(memory.ErrUserAlreadyExists)
			util.WriteJSONError(w, http.StatusBadRequest, err.Error())
			return
		}

		util.WriteJSONResponse(w, http.StatusOK, u)
	}
}

func (ur *UserRoutes) deleteUser() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		uidStr := chi.URLParam(r, "uid")

		if err := ur.userService.DeleteUser(uidStr); err != nil {
			util.WriteJSONError(w, http.StatusNotFound, err.Error())
			return
		}

		w.Write([]byte("User Deleted."))

	}
}
