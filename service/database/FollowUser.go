package database

// db function to execute the follow
func (db *appdbimpl) FollowUser(username string, userToFollow string) error {
	_, err := db.c.Exec("INSERT INTO followers(follower, followed) VALUES (?, ?)", username, userToFollow)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) CheckForFollow(u string, s string) (bool, error) {

	var counter int
	// query the db
	err := db.c.QueryRow("SELECT count(*) FROM followers WHERE follower = ? AND followed = ?",
		u, s).Scan(&counter)

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
