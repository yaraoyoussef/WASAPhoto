package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

// DONE

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// validation of user
	username := ps.ByName("username")
	userReq := extractBearer(r.Header.Get("Authorization"))
	valid := validateUser(username, userReq)
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}

	// initialize photo structure
	post := Photo{
		Owner:       username,
		DateAndTime: time.Now().UTC(),
	}

	// parse request body and extract raw picture
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.WithError(err).Error("error loading content")
		return
	}

	r.Body = io.NopCloser(bytes.NewBuffer(data))

	// generate photo's unique identifier
	id, err := rt.db.UploadPhoto(post.ToDatabase())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("Error occured in database")
		return
	}
	// convert to use in storing image in user's directory
	photoId := strconv.FormatInt(id, 10)

	// create dir for user to save images
	path := filepath.Join("userFolder", username, "photos")
	// create file to store content
	res, err := os.Create(filepath.Join(path, photoId))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("error creating photo file")
		return
	}
	// copy content to empty created file
	_, err = io.Copy(res, r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("error saving picture")
		return
	}

	// close created file
	res.Close()

	// return result to user
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(Photo{
		Comments:    nil,
		Likes:       nil,
		Owner:       post.Owner,
		DateAndTime: post.DateAndTime,
		ID:          int(id),
	})
}
