package api

import (
	"net/http"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

// function to remove comment from photo
func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	reqUser := extractBearer(r.Header.Get("Authorization"))

	// if user is not logged in
	if reqUser == "" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// check if user was banned by owner
	banned, err := rt.db.CheckForBan(ps.ByName("username"), reqUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("error while executing query")
		return
	}
	if banned {
		w.WriteHeader(http.StatusForbidden)
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

	// convert commentId from string to int
	commentId, err := strconv.ParseInt(ps.ByName("commentId"), 10, 64)
	// handle error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("error converting the comment's id")
		return
	}

	// check if owner of post is removing the comment
	if ps.ByName("username") == reqUser {
		err = rt.db.UncommentPhotoOwner(photoId, commentId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ctx.Logger.WithError(err).Error("error while executing db query")
			return
		}

		w.WriteHeader(http.StatusNoContent)
		return
	}

	// user is removing his own comment
	err = rt.db.UncommentPhoto(reqUser, photoId, commentId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("error while executing db query")
		return
	}

	// response
	w.WriteHeader(http.StatusNoContent)
}
