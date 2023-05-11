package api

import (
	"errors"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// extract both usernames from query
	currentUser := ps.ByName("id")
	userToFollow := ps.ByName("otherUserId")

	// extract from header
	userReq := extractBearer(r.Header.Get("Authorization"))
	// validation
	valid := validateUser(currentUser, userReq)
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}

	// user cannot follow himself
	if currentUser == userToFollow {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// check for ban in database
	banned, err := rt.db.CheckForBan(userToFollow, currentUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// if current user is banned, he cannot perform the follow
	if banned {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// call db function to update the database
	err = rt.db.FollowUser(currentUser, userToFollow)
	if errors.Is(err, database.ErrUserDoesNotExist) {
		// if user to follow does not exist, return 4XX error
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		// handle any other errors
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// response
	w.WriteHeader(http.StatusNoContent)
}
