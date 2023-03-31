package database

func (db *appdbimpl) UploadPhoto(p Photo) (Photo, error) {
	// add photo to list of photos; fetch the list, append it and push it back
	res, err := db.c.Exec(`UPDATE Profile SET (username, photos, followers, following, posts) VALUES (NULL, ?, ?, ?)`
				p.Username )
	if err != nil {
		return p, err
	}
	lastInsertId, err := res.LastInsertId()
	if err != nil {
		return p, err
	}
	p.ID = uint64(lastInsertId)
	return p, nil
}
