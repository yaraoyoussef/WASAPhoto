package database

// db function that retrieves the profile of a user given its username
func (db *appdbimpl) GetProfile(user User, reqUser User) (Profile, error) {
	var profile Profile
	following, err := db.GetFollowing(user.ID)
	if err != nil {
		return profile, err
	}
	followers, err := db.GetFollowers(user.ID)
	if err != nil {
		return profile, err
	}
	photos, err := db.GetPhotos(reqUser.ID, user.ID)
	if err != nil {
		return profile, err
	}
	username, err := db.GetUsername(user.ID)
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
func (db *appdbimpl) GetFollowing(user string) ([]User, error) {

	rows, err := db.c.Query("SELECT followed FROM followers WHERE follower = ?", user)
	if err != nil {
		return nil, err
	}
	// Wait for the function to finish before closing rows.
	defer func() { _ = rows.Close() }()

	// Read all the users appending the list
	var following []User
	for rows.Next() {
		var followed string
		err = rows.Scan(&followed)
		if err != nil {
			return nil, err
		}
		user := User{ID: followed}
		following = append(following, user)
	}

	if rows.Err() != nil {
		return nil, err
	}
	return following, nil
}

// db function that retrieves the list of users following the current user
func (db *appdbimpl) GetFollowers(user string) ([]User, error) {

	rows, err := db.c.Query("SELECT follower FROM followers WHERE followed = ?", user)
	if err != nil {
		return nil, err
	}
	// Wait for the function to finish before closing rows.
	defer func() { _ = rows.Close() }()

	// Read all the users appending the list
	var followers []User
	for rows.Next() {
		var follower string
		err = rows.Scan(&follower)
		user := User{ID: follower}
		if err != nil {
			return nil, err
		}
		followers = append(followers, user)
	}

	if rows.Err() != nil {
		return nil, err
	}
	return followers, nil
}

// db function that retrieves the list of photos of a user
func (db *appdbimpl) GetPhotos(reqUser string, user string) ([]Photo, error) {

	rows, err := db.c.Query("SELECT * FROM photos WHERE userId = ? ORDER BY dateAndTime DESC", user)
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

	rows, err := db.c.Query("SELECT * FROM comments WHERE photoId = ? AND userId NOT IN (SELECT uBannedId FROM banned WHERE userId = ? OR userId = ?) "+
		"AND userId NOT IN (SELECT userId FROM banned WHERE userId = ?)",
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
		var photoId int64
		err = rows.Scan(&comment.CommentId, &photoId, &comment.UserId, &comment.Comment)
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
func (db *appdbimpl) GetLikes(reqUser string, user string, photo int64) ([]User, error) {

	rows, err := db.c.Query("SELECT userId FROM likes WHERE photoId = ? AND userId NOT IN (SELECT uBannedId FROM banned WHERE uBannedId = ? OR userId = ?) "+
		"AND userId NOT IN (SELECT userId FROM banned WHERE uBannedId = ?)",
		photo, reqUser, user, reqUser)
	if err != nil {
		return nil, err
	}
	// Wait for the function to finish before closing rows.
	defer func() { _ = rows.Close() }()

	// Read all the users in the list
	var likes []User
	for rows.Next() {
		var lUser string
		err = rows.Scan(&lUser)
		user := User{ID: lUser}
		if err != nil {
			return nil, err
		}

		likes = append(likes, user)
	}

	if rows.Err() != nil {
		return nil, err
	}
	return likes, nil

}

// db function to get username
func (db *appdbimpl) GetUsername(user string) (string, error) {
	var username string
	err := db.c.QueryRow("SELECT username FROM users WHERE id =?", user).Scan(&username)
	if err != nil {
		return "", err
	}
	return username, nil

}
