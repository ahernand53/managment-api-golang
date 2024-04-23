package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	store Store
}

func NewUserService(s Store) *UserService {
	return &UserService{store: s}
}

func (s *UserService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/users/register", s.handleUserRegister).Methods("POST")
	// r.HandleFunc("/users/login", s.handleUserLogin).Methods("POST")
}

func (s *UserService) handleUserRegister(w http.ResponseWriter, r *http.Request) {
	// create user
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Error: "Error reading request body"})
		return
	}

	defer r.Body.Close()

	var payload *User
	err = json.Unmarshal(body, &payload)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid request payload"})
		return
	}

	if err := validateUserPayload(payload); err != nil {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	hashedPassword, err := HashPassword(payload.Password)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Error: "Error hashing password"})
		return
	}

	payload.Password = hashedPassword

	u, err := s.store.CreateUser(payload)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	token, err := createAndSetAuthCookie(u.ID, w)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Error: "Error creating session"})
		return
	}

	WriteJSON(w, http.StatusCreated, token)

}

func createAndSetAuthCookie(id int64, w http.ResponseWriter) (string, error) {
	secret := []byte(Envs.JWTSecret)
	token, err := CreateJWT(secret, id)
	if err != nil {
		return "", err
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "Authorization",
		Value: token,
	})

	return token, nil
}

func HashPassword(s string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func validateUserPayload(u *User) error {
	if u.Email == "" {
		return errors.New("email is required")
	}

	if u.Password == "" {
		return errors.New("password is required")
	}

	if u.FirstName == "" {
		return errors.New("first name is required")
	}

	if u.LastName == "" {
		return errors.New("last name is required")
	}

	return nil
}
