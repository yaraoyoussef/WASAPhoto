package api

import (
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// extract from header
	userReq := extractBearer(r.Header.Get("Authorization"))
	// validation
	valid := validateUser(ps.ByName("username"), userReq)
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}

	//TODO implement method that returns a stream of pictures for the current user's feed
	// get the user's followings, retrieve their posts, show to user
}
