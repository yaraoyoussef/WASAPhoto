package api

// DONE

import (
	"errors"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// extract both usernames from path
	currentUser := ps.ByName("id")
	userToBan := ps.ByName("otherUserId")

	// extract from header
	userReq := extractBearer(r.Header.Get("Authorization"))
	// validation
	valid := validateUser(currentUser, userReq)
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}

	// check if user is trying to ban himself
	if currentUser == userToBan {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// update database according to this ban
	err := rt.db.BanUser(currentUser, userToBan)

	// error handling
	if errors.Is(err, database.ErrUserDoesNotExist) {
		// if user to ban does not exist, return 4XX error
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		// handle any other errors
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// banning also means that current user will unfollow other user (if exists)
	// check for follow
	followed, err := rt.db.CheckForFollow(currentUser, userToBan)
	// error handling
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// if current user followed the banned user, unfollow him
	if followed {
		err = rt.db.UnfollowUser(currentUser, userToBan)
		// error handling
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	// banned user cannot follow user anymore
	// check for follow
	followed, err = rt.db.CheckForFollow(userToBan, currentUser)
	// error handling
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// if followed, unfollow
	if followed {
		err = rt.db.UnfollowUser(userToBan, currentUser)
		// handling error
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	// response
	w.WriteHeader(http.StatusNoContent)
}
