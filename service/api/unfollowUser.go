package api

import (
	"errors"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// get both current user and user to unfollow from query
	id := ps.ByName("id")
	otherUserId := ps.ByName("otherUserId")

	// extract from header
	userReq := extractBearer(r.Header.Get("Authorization"))
	// validation
	valid := validateUser(id, userReq)
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}

	// do operation on database
	err := rt.db.UnfollowUser(id, otherUserId)

	// check for errors and handle them
	if errors.Is(err, database.ErrUserDoesNotExist) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// action performed
	w.WriteHeader(http.StatusNoContent)
}
