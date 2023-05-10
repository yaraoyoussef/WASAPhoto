package database

// function that removes a like from database
func (db *appdbimpl) UnlikePhoto(photoId int64, userId string) error {

	_, err := db.c.Exec("DELETE FROM likes WHERE (photoId = ? AND userId = ?)", photoId, userId)
	if err != nil {
		return err
	}
	return nil
}
