package api

// DONE

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

// function that handles user's Login
func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// read content of the user from request body
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		// body was not parsable, so it gets rejected
		w.WriteHeader(http.StatusBadRequest)
		return
		// validation of user's structure content
	} else if !user.IsValid(user.ID) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// send user info to db
	err = rt.db.Login(user.ToDatabase())
	if err != nil {
		// user already exists in db
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(User{
			ID: user.ID,
		})
		if err != nil {
			// handle error on our side. Log it and send 500 to user
			// 500 code error should be in api operation responses??
			ctx.Logger.WithError(err).Error("cannot login user")
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	// create user's directories
	err = createUserDir(user.ID, ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("couldn't create user's folder")
		return
	}

	// send output to user
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(User{
		ID: user.ID,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("cannot create response json file")
	}
}

// function that creates new sub-dir for user
func createUserDir(id string, ctx reqcontext.RequestContext) error {
	// create path for each user
	path := filepath.Join("userFolder", id)
	// create sub-path in path for pictures of each user
	err := os.MkdirAll(filepath.Join(path, "photos"), os.ModePerm)
	// handle errors
	if err != nil {
		ctx.Logger.WithError(err).Error("user's directory couldn't be created")
		return err
	}
	return nil
}
