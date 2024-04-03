package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/security"
	"api/src/templates"
	"encoding/json"
	"io"
	"net/http"
)

// Login authenticates a user on the API
func Login(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		templates.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		templates.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		templates.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepository := repositories.NewUserRepository(db)
	userSavedOnDataBase, err := userRepository.SearchByEmail(user.Email)
	if err != nil {
		templates.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.ValidatePassword(userSavedOnDataBase.Password, user.Password); err != nil {
		templates.Error(w, http.StatusUnauthorized, err)
		return
	}

	token, _ := authentication.CreateToken(userSavedOnDataBase.ID)
	w.Write([]byte(token))
}
