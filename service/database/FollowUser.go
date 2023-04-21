package database

// db function to execute the follow
func (db *appdbimpl) FollowUser(username string, userToFollow string) error {
	_, err := db.c.Exec("INSERT INTO followers(follower, followed) VALUES (?, ?)", username, userToFollow)
	if err != nil {
		return err
	}
	return nil
}
