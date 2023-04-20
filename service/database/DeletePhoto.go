package database

func (db *appdbimpl) DeletePhoto(owner string, photoId int64) error {

	// query to execute on database
	_, err := db.c.Exec("DELETE FROM photos WHERE id_user = ? AND id_photo = ? ", owner, photoId)
	// check for errors
	if err != nil {
		return err
	}
	return nil
}
