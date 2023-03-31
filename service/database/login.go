package database

import (
	"strconv"
)

func (db *appdbimpl) Login(p User) (User, error) {
	res, err := db.c.Exec(`INSERT INTO users (id, username) VALUES (NULL, ?)`, p.Username)
	if err != nil {
		return p, err
	}
	// last insert id? non credo perche lo giestiamo usando username... but what about identifier generated
	id, err := res.LastInsertId()
	if err != nil {
		return p, err
	}

	p.ID = strconv.Itoa(int(id))
	return p, nil
}
