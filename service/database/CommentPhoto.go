package database

// Database function to add a comment to a picture
func (db *appdbimpl) CommentPhoto(pId int64, user string, comment Comment) (int64, error) {
	res, err := db.c.Exec("INSERT INTO comments (photoId, username, comment) VALUES (?, ?, ?)",
		pId, user, comment.Comment)

	// check for error while executing query
	if err != nil {
		return -1, err
	}

	// generate comment id
	commentId, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	return commentId, nil

}
