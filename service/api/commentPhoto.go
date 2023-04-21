package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var comment Comment

	// extract photo owner and the commenter
	photoOwner := ps.ByName("username")
	commenter := extractBearer(r.Header.Get("Authorization"))

	// if commenter is not logged in
	if commenter == "" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// check if commenter was banned by owner
	banned, err := rt.db.CheckForBan(photoOwner, commenter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("error while executing query")
		return
	}
	if banned {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// extract comment of user into comment structure
	err = json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.WithError(err).Error("could not parse request body")
		return
	}

	// check if max length for comment has been exceeded
	if len(comment.Comment) > 500 {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.WithError(err).Error("comment exceeds length limit (500 characters)")
		return
	}

	// convert photoId from string back to int
	photoId, err := strconv.ParseInt(ps.ByName("photoId"), 10, 64)
	// handle error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("error converting the picture's id")
		return
	}

	// call db fucntion for comment creating
	commentId, err := rt.db.CommentPhoto(photoId, commenter, comment.ToDatabase())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("could not post comment")
		return
	}

	// return to user
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(Comment{
		CommentId: commentId,
		Username:  commenter,
		Comment:   comment.Comment,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
