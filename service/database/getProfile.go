package database

// db function that retrieves the profile of a user given its username
func (db *appdbimpl) GetProfile(username string, reqUser string) (Profile, error) {
	var profile Profile
	following, err := db.GetFollowing(username)
	if err != nil {
		return profile, err
	}
	followers, err := db.GetFollowers(username)
	if err != nil {
		return profile, err
	}
	photos, err := db.GetPhotos(reqUser, username)
	if err != nil {
		return profile, err
	}
	posts := len(photos)

	profile = Profile{
		Username:  username,
		Followers: followers,
		Following: following,
		Photos:    photos,
		Posts:     posts,
	}

	return profile, nil
}

// db function that retrieves the list of users followed by the user
func (db *appdbimpl) GetFollowing(user string) ([]string, error) {

	rows, err := db.c.Query("SELECT followed FROM followers WHERE follower = ?", user)
	if err != nil {
		return nil, err
	}
	// Wait for the function to finish before closing rows.
	defer func() { _ = rows.Close() }()

	// Read all the users appending the list
	var following []string
	for rows.Next() {
		var followed string
		err = rows.Scan(&followed)
		if err != nil {
			return nil, err
		}
		following = append(following, followed)
	}

	if rows.Err() != nil {
		return nil, err
	}
	return following, nil
}

// db function that retrieves the list of users following the current user
func (db *appdbimpl) GetFollowers(user string) ([]string, error) {

	rows, err := db.c.Query("SELECT follower FROM followers WHERE followed = ?", user)
	if err != nil {
		return nil, err
	}
	// Wait for the function to finish before closing rows.
	defer func() { _ = rows.Close() }()

	// Read all the users appending the list
	var followers []string
	for rows.Next() {
		var follower string
		err = rows.Scan(&follower)
		if err != nil {
			return nil, err
		}
		followers = append(followers, follower)
	}

	if rows.Err() != nil {
		return nil, err
	}
	return followers, nil
}

// db function that retrieves the list of photos of a user
func (db *appdbimpl) GetPhotos(reqUser string, user string) ([]Photo, error) {

	rows, err := db.c.Query("SELECT * FROM photos WHERE user = ? ORDER BY dateAndTime DESC", user)
	if err != nil {
		return nil, err
	}
	// Wait for the function to finish before closing rows
	defer func() { _ = rows.Close() }()

	// Read all the photos in the list
	var photos []Photo
	for rows.Next() {
		var photo Photo
		err = rows.Scan(&photo.ID, &photo.Owner, &photo.DateAndTime)
		if err != nil {
			return nil, err
		}

		comments, err := db.GetComments(reqUser, user, int64(photo.ID))
		if err != nil {
			return nil, err
		}
		photo.Comments = comments

		likes, err := db.GetLikes(reqUser, user, int64(photo.ID))
		if err != nil {
			return nil, err
		}
		photo.Likes = likes

		photos = append(photos, photo)
	}

	if rows.Err() != nil {
		return nil, err
	}
	return photos, nil
}

// Database function that retrieves the list of comments of a photo
func (db *appdbimpl) GetComments(reqUser string, user string, photo int64) ([]Comment, error) {

	rows, err := db.c.Query("SELECT * FROM comments WHERE photoId = ? AND username NOT IN (SELECT uBanned FROM banned WHERE user = ? OR user = ?) "+
		"AND username NOT IN (SELECT user FROM banned WHERE user = ?)",
		photo, reqUser, user, reqUser)
	if err != nil {
		return nil, err
	}

	// Wait for the function to finish before closing rows
	defer func() { _ = rows.Close() }()

	// Read all the comments
	var comments []Comment
	for rows.Next() {
		var comment Comment
		err = rows.Scan(&comment.CommentId, &comment.Username, &comment.Comment)
		if err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	if rows.Err() != nil {
		return nil, err
	}
	return comments, nil
}

// Database function that retrieves the list of users that liked a photo
func (db *appdbimpl) GetLikes(reqUser string, user string, photo int64) ([]string, error) {

	rows, err := db.c.Query("SELECT username FROM likes WHERE photoId = ? AND username NOT IN (SELECT uBanned FROM banned WHERE uBanner = ? OR user = ?) "+
		"AND username NOT IN (SELECT user FROM banned WHERE uBanned = ?)",
		photo, reqUser, user, reqUser)
	if err != nil {
		return nil, err
	}
	// Wait for the function to finish before closing rows.
	defer func() { _ = rows.Close() }()

	// Read all the users in the list
	var likes []string
	for rows.Next() {
		var lUser string
		err = rows.Scan(&lUser)
		if err != nil {
			return nil, err
		}

		likes = append(likes, lUser)
	}

	if rows.Err() != nil {
		return nil, err
	}
	return likes, nil

}
