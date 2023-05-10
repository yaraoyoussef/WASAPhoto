package database

// db function to change the username
func (db *appdbimpl) SetUsername(user User, username Username) error {

	// modify username of current user
	res, err := db.c.Exec(`UPDATE users SET username=? WHERE id=?`, username.Username, user.ID)
	if err != nil {
		return err
	}

	aff, err := res.RowsAffected()
	if err != nil {
		return err
	} else if aff == 0 {
		// row wasnt modified
		return ErrCouldNotModify
	}
	return nil
}
