package database

func (db *appdbimpl) FollowUser(username string, userToFollow string) error {
	var p Profile
	var p1 Profile
	err := db.c.QueryRow("SELECT * FROM profiles WHERE username = ?", userToFollow).Scan(&p)
	// check if user does not exist error
	if err != nil {
		return ErrUserDoesNotExist
	}
	// save current user profile in p1 variable
	err = db.c.QueryRow("SELECT * FROM profiles WHERE username = ?", username).Scan(&p1)
	if err != nil {
		return err
	}
	// add current user to followers list of otherUser
	_, err = db.c.Exec("UPDATE profiles SET followers=? WHERE username=?", append(p.Followers, username), p.Username)
	if err != nil {
		return err
	}

	// add user to follow to following's list of current user
	_, err = db.c.Exec("UPDATE profiles SET following=? WHERE username=?", append(p1.Following, p.Username), p1.Username)
	if err != nil {
		return err
	}
	return nil
}
