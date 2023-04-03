package api

import (
	"time"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
)

type User struct {
	ID       string `json:"ID"`
	Username string `json:"username"`
}

type Username struct {
	Username string `json:"username"`
}

type Profile struct {
	Username string `json:"username"`
	//  followers/following should be database.user
	Photos    []database.Photo `json:"photos"`
	Followers []string         `json:"followers"`
	Following []string         `json:"following"`
	Posts     int              `json:"posts"`
}

type Stream struct {
	Photos []database.Photo `json:"data"`
}

type Photo struct {
	ID          int                `json:"photoId"`
	Owner       string             `json:"owner"`
	Likes       []database.User    `json:"likes"`
	Comments    []database.Comment `json:"comments"`
	DateAndTime time.Time          `json:"dateAndTime"`
	// in api doc i have the picture in here too, but idk how to put it here
	// maybe will modify api
}

type Comment struct {
	CommentId int64  `json:"commentId"`
	Comment   string `json:"comment"`
	Username  string `json:"username"`
}

func (p *User) ToDatabase() database.User {
	return database.User{
		ID:       p.ID,
		Username: p.Username,
	}
}

func (p *Profile) ToDatabase() database.Profile {
	return database.Profile{
		Username:  p.Username,
		Photos:    p.Photos,
		Followers: p.Followers,
		Following: p.Following,
		Posts:     p.Posts,
	}
}

func (p *Profile) FromDatabase(profile database.Profile) {
	p.Username = profile.Username
	p.Photos = profile.Photos
	p.Followers = profile.Followers
	p.Following = profile.Following
	p.Posts = profile.Posts
}

func (p *Photo) ToDatabase() database.Photo {
	return database.Photo{
		ID:          p.ID,
		Owner:       p.Owner,
		Likes:       p.Likes,
		Comments:    p.Comments,
		DateAndTime: p.DateAndTime,
	}
}
