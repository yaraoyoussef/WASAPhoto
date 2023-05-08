package api

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// get photo id and username
	username := ps.ByName("username")
	photo := ps.ByName("photoId")

	// validation
	valid := validateUser(username, extractBearer(r.Header.Get("Authorization")))
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}

	// convert photoId from string back to int
	photoId, err := strconv.ParseInt(photo, 10, 64)
	// handle error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("error converting the picture's id")
		return
	}

	// call to db function to delete picture
	err = rt.db.DeletePhoto(username, photoId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("an error occured, could not delete post")
		return
	}

	// eliminate file of the picture in question
	path := filepath.Join("userFolder", username, "photos")
	err = os.Remove(filepath.Join(path, photo))
	if err != nil {
		ctx.Logger.WithError(err).Error("photo does not exist")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// response
	w.WriteHeader(http.StatusNoContent)

}
