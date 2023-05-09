package database

// db function to loin
func (db *appdbimpl) Login(p User) (string, error) {

	// insert into db using query
	_, err := db.c.Exec(`INSERT INTO users (id, username) VALUES (?, ?)`, p.Username, p.Username)

	if err != nil {
		print(err.Error())
		return "", err
	}

	err = db.c.QueryRow(`SELECT id FROM users WHERE username = ?`, p.Username).Scan(&p.ID)
	if err != nil {
		return "", err
	}

	return p.ID, nil
}
