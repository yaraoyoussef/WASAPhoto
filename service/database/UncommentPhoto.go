package database

// db function to uncomment a photo given the photo id and the comment id
func (db *appdbimpl) UncommentPhotoOwner(pId int64, cId int64) error {
	_, err := db.c.Exec("DELETE FROM comments WHERE (photoId = ? AND commentId = ?)", pId, cId)
	if err != nil {
		return err
	}

	return nil
}

// db function to uncomment a photo given the photo id, comment id and username
func (db *appdbimpl) UncommentPhoto(userId string, pId int64, cId int64) error {
	_, err := db.c.Exec("DELETE FROM comments WHERE (photoId = ? AND userId = ? AND commentId = ?)",
		pId, userId, cId)
	if err != nil {
		return err
	}

	return nil
}
