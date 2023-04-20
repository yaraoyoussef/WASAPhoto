package database

// function that removes a like from database
func (db *appdbimpl) UnlikePhoto(photoId int64, username string) error {

	_, err := db.c.Exec("DELETE FROM likes WHERE (photoId = ? AND username = ?)", photoId, username)
	if err != nil {
		return err
	}
	return nil
}
