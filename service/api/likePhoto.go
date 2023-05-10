package api

import (
	"net/http"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

// function to like post of another user
func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userId := ps.ByName("id")
	userLiked := ps.ByName("otherUserId")

	// user cannot like his own picture
	if userLiked == userId {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// validation (userLiked) bcz to be able to put a like, user has to be logged in
	valid := validateUser(userLiked, extractBearer(r.Header.Get("Authorization")))
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}

	// check if user who wants to like was banned by the user owner of the post
	banned, err := rt.db.CheckForBan(userId, userLiked)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("an error occured in database")
		return
	}
	if banned {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// convert photoId to int64
	photoId, err := strconv.ParseInt(ps.ByName("photoId"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("an error occured during conversion")
		return
	}

	// put like (by inserting it into database)
	err = rt.db.LikePhoto(photoId, userLiked)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// response
	w.WriteHeader(http.StatusNoContent)

}
