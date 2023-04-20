package database

func (db *appdbimpl) LikePhoto(photoId int64, username string) error {
	_, err := db.c.Exec("INSERT INTO likes (photoId, username) VALUES (?, ?)", photoId, username)
	if err != nil {
		return err
	}
	return nil
}
