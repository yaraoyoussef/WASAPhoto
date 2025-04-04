package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// extract from header
	userReq := extractBearer(r.Header.Get("Authorization"))
	// validation
	valid := validateUser(ps.ByName("id"), userReq)
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}

	// get the list of users that current user follows
	followings, err := rt.db.GetFollowing(userReq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// get photos of each of the followings
	var photos []database.Photo
	for _, following := range followings {
		followingPhotos, err := rt.db.GetPhotos(userReq, following.ID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// limit number of photos per following user in home page of current user
		for i, photo := range followingPhotos {
			if i >= database.PPU {
				break
			}
			photos = append(photos, photo)
		}
	}

	// return to user
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(photos)
}
