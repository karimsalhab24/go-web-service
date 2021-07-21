package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"github.com/karimsalhab24/go-web-service/models"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

//we are binding this method to a user controller, this is the difference
// between a function and a method
//IMP: since this file uses this method, it is automatically assumed by Go to be
// implementing the interface Handler which holds this method
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// path contains users so we are operating on multiple users
	if r.URL.Path == "/users" {
		switch r.Method {
		//getting users
		case http.MethodGet:
			uc.getAll(w, r)
		//adding users
		case http.MethodPost:
			uc.post(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		//check if the path matches the user id regex
		matches := uc.userIDPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
		}

		//convert String response to a numerical data
		id, err := strconv.Atoi(matches[1])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}
		switch r.Method {
		//getting user
		case http.MethodGet:
			uc.get(id, w)
		//adding user
		case http.MethodPut:
			uc.put(id, w, r)
		//deleting user
		case http.MethodDelete:
			uc.delete(id, w)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}

	}
	w.Write([]byte("Hello from the User Controller"))
}

func (uc *userController) getAll(w http.ResponseWriter, r *http.Request) {
	encoreResponseAsJSON(models.GetUsers(), w)
}

func (uc *userController) get(id int, w http.ResponseWriter) {
	u, err := models.GetUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encoreResponseAsJSON(u, w)
}

func (uc *userController) post(w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse user object"))
		return
	}
	u, err = models.AddUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse user object"))
		return
	}
}

func (uc *userController) put(id int, w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse user object"))
		return
	}
	if id != u.ID {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("ID of submitted user must match ID in URL"))
		return
	}
	encoreResponseAsJSON(u, w)
}

func (uc *userController) delete(id int, w http.ResponseWriter) {
	err := models.RemoveUserById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

//parse the coming request and get the user requested if exists
func (uc *userController) parseRequest(r *http.Request) (models.User, error) {
	dec := json.NewDecoder(r.Body)
	var u models.User
	err := dec.Decode(&u)
	if err != nil {
		return models.User{}, err
	}
	return u, nil
}

//go convention: start a func name with "new" to specify that it is a constructor function
func newUserController() *userController {
	// we can use the address of (&) on a struct instead of a var name
	// but we cannot do for example: &42 (literal variable)
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(d+)/?`),
	}
}
