package database

func (db *appdbimpl) GetMyStream(username string) ([]Photo, error) {
	rows, err := db.c.Query(`SELECT * FROM photos WHERE user IN (SELECT followed FROM followers WHERE follower = ?)
	ORDER BY dateAndTime DESC`, username)
	if err != nil {
		return nil, err
	}

	// wait for the funtion to finish before closing rows
	defer func() { _ = rows.Close() }()

	// read all users into list
	var res []Photo
	for rows.Next() {
		var photo Photo
		err = rows.Scan(&photo.ID, &photo.Owner, &photo.DateAndTime)
		if err != nil {
			return nil, err
		}
		res = append(res, photo)
	}

	if rows.Err() != nil {
		return nil, err
	}
	return res, nil

}
