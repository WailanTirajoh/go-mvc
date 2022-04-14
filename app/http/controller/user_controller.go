package controller

import (
	"encoding/json"
	"net/http"

	"github.com/WailanTirajoh/go-simple-clean-architecture/go-simple-clean-architecture/app/helper"
	"github.com/WailanTirajoh/go-simple-clean-architecture/go-simple-clean-architecture/app/http/service"
	"github.com/gorilla/mux"
)

func NewUserController(userService *service.UserService) UserController {
	return UserController{
		UserService: *userService,
	}
}

type UserController struct {
	UserService service.UserService
}

func (userController *UserController) Route(router *mux.Router) {
	router.HandleFunc("/users", userController.Index).Methods("GET")
	router.HandleFunc("/users/{id}", userController.Show).Methods("GET")
	router.HandleFunc("/users", userController.Store).Methods("POST")
	router.HandleFunc("/users/{id}", userController.Update).Methods("PUT")
	router.HandleFunc("/users/{id}", userController.Destroy).Methods("DELETE")
}

func (userController *UserController) Index(writter http.ResponseWriter, request *http.Request) {
	writter.Header().Set("Content-Type", "application/json")

	users := userController.UserService.GetUsers()

	json.NewEncoder(writter).Encode(helper.SuccessResponse(users))
}

func (userController *UserController) Show(writter http.ResponseWriter, request *http.Request) {
	writter.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)
	user, err := userController.UserService.GetUser(params["id"])

	if err != nil {
		writter.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writter).Encode(helper.ErrorResponse(err.Error()))
		return
	}

	json.NewEncoder(writter).Encode(helper.SuccessResponse(user))
}

func (userController *UserController) Store(writter http.ResponseWriter, request *http.Request) {
	writter.Header().Set("Content-Type", "application/json")

	user := userController.UserService.StoreUser(request.Body)

	json.NewEncoder(writter).Encode(helper.SuccessResponse(user))
}

func (userController *UserController) Update(writter http.ResponseWriter, request *http.Request) {
	writter.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)
	user, err := userController.UserService.UpdateUser(params["id"], request.Body)

	if err != nil {
		writter.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writter).Encode(helper.ErrorResponse(err.Error()))
		return
	}

	json.NewEncoder(writter).Encode(helper.SuccessResponse(user))
}

func (userController *UserController) Destroy(writter http.ResponseWriter, request *http.Request) {
	writter.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)
	message, err := userController.UserService.DeleteUser(params["id"])

	if err != nil {
		writter.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writter).Encode(helper.ErrorResponse(err.Error()))
		return
	}

	json.NewEncoder(writter).Encode(helper.SuccessResponse(message))
}
