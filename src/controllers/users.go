package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/security"
	"api/src/templates"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CreateUser creates a new user on database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		templates.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		templates.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare(models.CREATE); err != nil {
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
	user.ID, err = userRepository.Create(user)
	if err != nil {
		templates.Error(w, http.StatusInternalServerError, err)
		return
	}

	templates.JSON(w, http.StatusCreated, user)
}

// FindUsers retrieve all users from the database
func FindUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	db, err := database.Connect()
	if err != nil {
		templates.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepository := repositories.NewUserRepository(db)
	users, err := userRepository.Search(nameOrNick)
	if err != nil {
		templates.Error(w, http.StatusInternalServerError, err)
		return
	}

	templates.JSON(w, http.StatusOK, users)
}

// FindUser retrieve one specified users from the database
func FindUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		templates.Error(w, http.StatusInternalServerError, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		templates.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepository := repositories.NewUserRepository(db)
	user, err := userRepository.SerachByID(userID)
	if err != nil {
		templates.Error(w, http.StatusInternalServerError, err)
		return
	}

	templates.JSON(w, http.StatusOK, user)
}

// UpdateUser updates one specified users from the database
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		templates.Error(w, http.StatusBadRequest, err)
		return
	}

	tokenUserID, err := authentication.ExtractUserID(r)
	if err != nil {
		templates.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userID != tokenUserID {
		templates.Error(w, http.StatusForbidden, errors.New("its not possible to update other users information, only your own"))
		return
	}

	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		templates.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		templates.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare(models.EDIT); err != nil {
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
	if err = userRepository.Update(userID, user); err != nil {
		templates.Error(w, http.StatusInternalServerError, err)
		return
	}

	templates.JSON(w, http.StatusNoContent, nil)
}

// DeleteUser deletes one specified users from the database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		templates.Error(w, http.StatusBadRequest, err)
		return
	}

	tokenUserID, err := authentication.ExtractUserID(r)
	if err != nil {
		templates.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userID != tokenUserID {
		templates.Error(w, http.StatusForbidden, errors.New("its not possible to delete other users, only your own"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		templates.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepository := repositories.NewUserRepository(db)
	if err = userRepository.Delete(userID); err != nil {
		templates.Error(w, http.StatusInternalServerError, err)
		return
	}

	templates.JSON(w, http.StatusNoContent, nil)
}

// FollowUser enables an user to follow another
func FollowUser(w http.ResponseWriter, r *http.Request) {
	followerID, err := authentication.ExtractUserID(r)
	if err != nil {
		templates.Error(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		templates.Error(w, http.StatusBadRequest, err)
		return
	}

	if userID == followerID {
		templates.Error(w, http.StatusForbidden, errors.New("its not possible to follow yourself"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		templates.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepository := repositories.NewUserRepository(db)
	if err = userRepository.FollowUser(userID, followerID); err != nil {
		templates.Error(w, http.StatusInternalServerError, err)
		return
	}

	templates.JSON(w, http.StatusNoContent, nil)
}

// UnFollowUser enables an user to follow another
func UnFollowUser(w http.ResponseWriter, r *http.Request) {
	followerID, err := authentication.ExtractUserID(r)
	if err != nil {
		templates.Error(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		templates.Error(w, http.StatusBadRequest, err)
		return
	}

	if userID == followerID {
		templates.Error(w, http.StatusForbidden, errors.New("its not possible to unfollow yourself"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		templates.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepository := repositories.NewUserRepository(db)
	if err = userRepository.UnFollowUser(userID, followerID); err != nil {
		templates.Error(w, http.StatusInternalServerError, err)
		return
	}

	templates.JSON(w, http.StatusNoContent, nil)
}

// SearchFollowers get all followers of an user
func SearchFollowers(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	userID, err := strconv.ParseUint(param["userId"], 10, 64)
	if err != nil {
		templates.Error(w, http.StatusBadRequest, err)
	}

	db, err := database.Connect()
	if err != nil {
		templates.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepository := repositories.NewUserRepository(db)
	followers, err := userRepository.SearchFollowers(userID)
	if err != nil {
		templates.Error(w, http.StatusInternalServerError, err)
		return
	}

	templates.JSON(w, http.StatusOK, followers)
}

// SearchFollowing get all users that an user follows
func SearchFollowing(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	userID, err := strconv.ParseUint(param["userId"], 10, 64)
	if err != nil {
		templates.Error(w, http.StatusBadRequest, err)
	}

	db, err := database.Connect()
	if err != nil {
		templates.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepository := repositories.NewUserRepository(db)
	following, err := userRepository.SearchFollowing(userID)
	if err != nil {
		templates.Error(w, http.StatusInternalServerError, err)
		return
	}

	templates.JSON(w, http.StatusOK, following)
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	tokenUserID, err := authentication.ExtractUserID(r)
	if err != nil {
		templates.Error(w, http.StatusBadRequest, err)
		return
	}

	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		templates.Error(w, http.StatusBadRequest, err)
		return
	}

	if userID != tokenUserID {
		templates.Error(w, http.StatusForbidden, errors.New("its not possible to update other users password"))
		return
	}

	bodyRequest, err := io.ReadAll(r.Body)

	var password models.Password
	if err = json.Unmarshal(bodyRequest, &password); err != nil {
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
	savedPasswordHash, err := userRepository.SearchPassword(userID)
	if err != nil {
		templates.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.ValidatePassword(savedPasswordHash, password.Current); err != nil {
		templates.Error(w, http.StatusUnauthorized, errors.New("the current password doesn't match users password"))
		return
	}

	hashedPassword, err := security.Hash(password.New)
	if err != nil {
		templates.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = userRepository.UpdatePassword(userID, string(hashedPassword)); err != nil {
		templates.Error(w, http.StatusInternalServerError, err)
		return
	}

	templates.JSON(w, http.StatusNoContent, nil)
}
