package database

// Database function that allows user to ban another user
func (db *appdbimpl) BanUser(cUser string, userTB string) error {
	_, err := db.c.Exec("INSERT INTO banned (username, uBanned) VALUES (?, ?)", cUser, userTB)
	if err != nil {
		return err
	}
	return nil
}

// checks if cUser banned bUser
func (db *appdbimpl) CheckForBan(cUser string, bUser string) (bool, error) {
	var counter int

	// query the db
	err := db.c.QueryRow("SELECT count(*) FROM banned WHERE username = ? AND uBanned = ?",
		cUser, bUser).Scan(&counter)

	// handle errors
	if err != nil {
		return true, err
	}

	// if counter is 1, then bUser is banned by cUser
	if counter == 1 {
		return true, nil
	}

	// else he is not
	return false, nil
}
