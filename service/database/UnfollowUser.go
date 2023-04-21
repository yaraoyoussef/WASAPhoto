package database

//DONE

// db function to execute the unfollow
func (db *appdbimpl) UnfollowUser(u string, s string) error {
	_, err := db.c.Exec("DELETE FROM followers WHERE (follower = ?, followed = ?) ", u, s)
	if err != nil {
		return err
	}
	return nil
}
