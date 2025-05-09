package database

// database function to delete a photo uploaded
func (db *appdbimpl) DeletePhoto(owner string, photoId int64) error {

	// query to execute on database
	_, err := db.c.Exec("DELETE FROM photos WHERE userId = ? AND photoId = ? ", owner, photoId)
	// check for errors
	if err != nil {
		return err
	}
	return nil
}
