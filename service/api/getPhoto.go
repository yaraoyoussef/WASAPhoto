package api

import (
	"net/http"
	"path/filepath"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	http.ServeFile(w, r, filepath.Join("userFolder", ps.ByName("id"), "photos", ps.ByName("photoId")))
}
