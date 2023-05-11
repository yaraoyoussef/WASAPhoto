package database

// db function to loin
func (db *appdbimpl) Login(p User) error {

	// insert into db using query
	_, err := db.c.Exec(`INSERT INTO users (id, username) VALUES (?, ?)`, p.ID, p.ID)

	if err != nil {
		return err
	}

	return nil
}
