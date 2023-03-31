package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

// TO COMPLETE

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var post Photo
	// parse request body and extract the picture from it
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// store picture in db
	dbPost, err := rt.db.UploadPhoto(post.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("Cannot upload post, please retry later")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// add picture to profile of user
	// to do that we need to update the db

}
