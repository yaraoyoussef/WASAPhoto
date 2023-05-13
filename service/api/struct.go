package api

import (
	"time"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
)

// structure that represents a user
type User struct {
	ID string `json:"ID"`
}

// structure that represents a username
type Username struct {
	Username string `json:"username"`
}

// structre that represents a profile
type Profile struct {
	Username  string           `json:"username"`
	Photos    []database.Photo `json:"photos"`
	Followers []database.User  `json:"followers"`
	Following []database.User  `json:"following"`
	Posts     int              `json:"posts"`
}

// structre that represents a photo
type Photo struct {
	ID          int                `json:"photoId"`
	Owner       string             `json:"owner"`
	Likes       []database.User    `json:"likes"`
	Comments    []database.Comment `json:"comments"`
	DateAndTime time.Time          `json:"dateAndTime"`
}

// structre that represents a comment
type Comment struct {
	CommentId int64  `json:"commentId"`
	Comment   string `json:"comment"`
	UserId    string `json:"userId"`
}

func (p User) ToDatabase() database.User {
	return database.User{
		ID: p.ID,
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

func (u Username) ToDatabase() database.Username {
	return database.Username{
		Username: u.Username,
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

func (c *Comment) ToDatabase() database.Comment {
	return database.Comment{
		CommentId: c.CommentId,
		Comment:   c.Comment,
		UserId:    c.UserId,
	}
}
