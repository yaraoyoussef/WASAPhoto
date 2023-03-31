package database

//DONE

func (db *appdbimpl) UnfollowUser(u string, s string) error {
	var toUnfollow Profile
	var p Profile

	// select user to unfollow from db and store it into var toUnfollow
	err := db.c.QueryRow("SELECT * FROM profiles WHERE username=?", s).Scan(&toUnfollow)
	if err != nil {
		return ErrUserDoesNotExist
	}

	// select current user from db and store it into var p
	err = db.c.QueryRow("SELECT * FROM profiles WHERE username=?", u).Scan(&p)
	if err != nil {
		return err
	}

	// search of name of current user in followers of the user to unfollow
	for i := range toUnfollow.Followers {
		if toUnfollow.Followers[i] == p.Username {
			// when found remove it
			toUnfollow.Followers[i] = toUnfollow.Followers[len(toUnfollow.Followers)-1]
			followers := toUnfollow.Followers[:len(toUnfollow.Followers)-1]
			_, err = db.c.Exec("UPDATE profiles SET followers=? WHERE username=?", followers, toUnfollow.Username)

			// handle any errors
			if err != nil {
				return err
			}
			break
		}
	}

	// search from name of user to unfollow in following of current user
	for j := range p.Following {
		if p.Following[j] == toUnfollow.Username {
			// when found remove it from the following's list
			p.Following[j] = p.Following[len(p.Following)-1]
			following := p.Following[:len(p.Following)-1]
			_, err = db.c.Exec("UPDATE profiles SET following=? WHERE username=?", following, p.Username)
			if err != nil {
				// handling errors
				return err
			}
			break
		}
	}

	return nil
}
