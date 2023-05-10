package database

// Database function to remove a ban
func (db *appdbimpl) UnbanUser(u string, s string) error {

	_, err := db.c.Exec("DELETE FROM banned WHERE (userId = ? AND uBannedId = ?)", u, s)
	if err != nil {
		return err
	}
	return nil
}
