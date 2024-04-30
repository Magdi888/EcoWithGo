package user

import (
	"fmt"
	"net/http"

	"github.com/Magdi888/EcoWithGo/types"
	"github.com/Magdi888/EcoWithGo/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store  types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}


func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handlerLogin).Methods("POST")
	router.HandleFunc("/register", h.handlerRegister).Methods("POST")

}

func (h *Handler) handlerLogin(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) handlerRegister(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, &payload); err  != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
// Check the  email	 is already in use or not
	_, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user  with this email already exists"))
		return
	}
	err =  h.store.CreateUser(types.User{
		FirstName:     payload.FirstName,
		LastName:      payload.LastName,
		Email:         payload.Email,
		Password:      payload.Password,
	})
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
		

}