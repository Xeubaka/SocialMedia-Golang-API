package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/templates"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreatePost creates a new post on the database
func CreatePost(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserID(r)
	if err != nil {
		templates.Error(w, http.StatusUnauthorized, err)
		return
	}

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		templates.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var post models.Post
	if err = json.Unmarshal(requestBody, &post); err != nil {
		templates.Error(w, http.StatusBadRequest, err)
		return
	}

	post.AuthorID = userID

	if err = post.Prepare(); err != nil {
		templates.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		templates.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	postRepository := repositories.NewPostRepository(db)
	post.ID, err = postRepository.CreatePost(post)
	if err != nil {
		templates.Error(w, http.StatusInternalServerError, err)
		return
	}

	templates.JSON(w, http.StatusCreated, post)
}

// FindPosts find all posts in the database
func FindPosts(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserID(r)
	if err != nil {
		templates.Error(w, http.StatusUnauthorized, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		templates.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	postRepository := repositories.NewPostRepository(db)
	posts, err := postRepository.Search(userID)
	if err != nil {
		templates.Error(w, http.StatusInternalServerError, err)
		return
	}

	templates.JSON(w, http.StatusOK, posts)
}

// FindPost find a single post in the database
func FindPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["postId"], 10, 64)
	if err != nil {
		templates.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		templates.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	postRepository := repositories.NewPostRepository(db)
	post, err := postRepository.SearchByID(postID)
	if err != nil {
		templates.Error(w, http.StatusInternalServerError, err)
		return
	}

	templates.JSON(w, http.StatusOK, post)
}

// UpdatePost update information of a single post
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserID(r)
	if err != nil {
		templates.Error(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["postId"], 10, 64)
	if err != nil {
		templates.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		templates.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	postRepository := repositories.NewPostRepository(db)
	postSavedOnDB, err := postRepository.SearchByID(postID)
	if err != nil {
		templates.Error(w, http.StatusInternalServerError, err)
		return
	}

	if postSavedOnDB.AuthorID != userID {
		templates.Error(w, http.StatusForbidden, errors.New("its not possible to update others user's posts"))
		return
	}

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		templates.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var post models.Post
	if err = json.Unmarshal(requestBody, &post); err != nil {
		templates.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = post.Prepare(); err != nil {
		templates.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = postRepository.UpdatePost(postID, post); err != nil {
		templates.Error(w, http.StatusInternalServerError, err)
		return
	}

	templates.JSON(w, http.StatusNoContent, nil)
}

// DeletePost delete a single post from the database
func DeletePost(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserID(r)
	if err != nil {
		templates.Error(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["postId"], 10, 64)
	if err != nil {
		templates.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		templates.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	postRepository := repositories.NewPostRepository(db)
	postSavedOnDB, err := postRepository.SearchByID(postID)
	if err != nil {
		templates.Error(w, http.StatusInternalServerError, err)
		return
	}

	if postSavedOnDB.AuthorID != userID {
		templates.Error(w, http.StatusForbidden, errors.New("its not possible to delete others user's posts"))
		return
	}

	if err = postRepository.DeletePost(postID); err != nil {
		templates.Error(w, http.StatusInternalServerError, err)
		return
	}

	templates.JSON(w, http.StatusNoContent, nil)
}

// SeachPostsByUser gets all posts from an user
func SeachPostsByUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		templates.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		templates.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	postRepository := repositories.NewPostRepository(db)
	posts, err := postRepository.SearchPostsByUser(userID)
	if err != nil {
		templates.Error(w, http.StatusInternalServerError, err)
		return
	}

	templates.JSON(w, http.StatusOK, posts)
}

// LikePost add a like on the post
func LikePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["postId"], 10, 64)
	if err != nil {
		templates.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		templates.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	postRepository := repositories.NewPostRepository(db)
	if err = postRepository.Like(postID); err != nil {
		templates.Error(w, http.StatusInternalServerError, err)
		return
	}

	templates.JSON(w, http.StatusNoContent, nil)
}

// UnLikePost removes a like from the post
func UnLikePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["postId"], 10, 64)
	if err != nil {
		templates.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		templates.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	postRepository := repositories.NewPostRepository(db)
	if err = postRepository.UnLike(postID); err != nil {
		templates.Error(w, http.StatusInternalServerError, err)
		return
	}

	templates.JSON(w, http.StatusNoContent, nil)
}
