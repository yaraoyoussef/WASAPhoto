package database

// function that retrieves the profile of a user given its username
func (db *appdbimpl) GetProfile(username string) (Profile, error) {
	var profile Profile
	err := db.c.QueryRow("SELECT * FROM profiles WHERE username=?", username).Scan(&profile)
	return profile, err
}
