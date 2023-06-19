package database

func (db *appdbimpl) SearchUsers(searching string, searchedFor string) ([]User, error) {
	rows, err := db.c.Query("SELECT * FROM users WHERE((id LIKE ?) OR (username LIKE ?)) AND id NOT IN (SELECT userId FROM banned WHERE ubannedId=?)", searchedFor, searchedFor, searching)
	if err != nil {
		return nil, err
	}

	defer func() { _ = rows.Close() }()

	var res []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID)
		if err != nil {
			return nil, err
		}
		res = append(res, user)
	}
	return res, nil
}
