package api

import (
	"net/http"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

// function to unlike photo
func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	photoOwner := ps.ByName("id")
	reqUser := ps.ByName("otherUserId")

	// validation
	valid := validateUser(reqUser, extractBearer(r.Header.Get("Authorization")))
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}

	// check if reqUser is banned by photo owner
	banned, err := rt.db.CheckForBan(photoOwner, reqUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("error occured in database")
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

	// removes like from database
	err = rt.db.UnlikePhoto(photoId, reqUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// response
	w.WriteHeader(http.StatusNoContent)
}
