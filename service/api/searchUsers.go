package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) searchUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// req user
	userReq := extractBearer(r.Header.Get("Authorization"))
	if userReq == "" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	searchedFor := r.URL.Query().Get("id")

	res, err := rt.db.SearchUsers(userReq, searchedFor)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("an error occured, please retry later")
		_ = json.NewEncoder(w).Encode([]User{})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(res)
}
