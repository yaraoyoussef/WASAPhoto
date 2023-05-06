package database

// db function to upload photo
func (db *appdbimpl) UploadPhoto(p Photo) (int64, error) {
	// add photo
	res, err := db.c.Exec("INSERT INTO photos (username, dateAndTime) VALUES (?,?)", p.Owner, p.DateAndTime)
	// check for errors
	if err != nil {
		return -1, err
	}
	// generate unique picture's id
	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}
	return id, nil

}
