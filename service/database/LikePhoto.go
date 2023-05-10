package database

// db function to like a photo
func (db *appdbimpl) LikePhoto(photoId int64, userId string) error {
	_, err := db.c.Exec("INSERT INTO likes (photoId, userId) VALUES (?, ?)", photoId, userId)
	if err != nil {
		return err
	}
	return nil
}
