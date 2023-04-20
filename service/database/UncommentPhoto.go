package database

func (db *appdbimpl) UncommentPhotoOwner(pId int64, cId int64) error {
	_, err := db.c.Exec("DELETE FROM comments WHERE (photoId = ? AND commentId = ?)", pId, cId)
	if err != nil {
		return err
	}

	return nil
}

func (db *appdbimpl) UncommentPhoto(username string, pId int64, cId int64) error {
	_, err := db.c.Exec("DELETE FROM comments WHERE (photoId = ? AND username = ? AND commentId = ?)",
		pId, username, cId)
	if err != nil {
		return err
	}

	return nil
}
