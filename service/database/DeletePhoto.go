package database

func (db *appdbimpl) DeletePhoto(owner string, photoId int64) error {

	// query to execute on database
	_, err := db.c.Exec("DELETE FROM photos WHERE user = ? AND photoId = ? ", owner, photoId)
	// check for errors
	if err != nil {
		return err
	}
	return nil
}
