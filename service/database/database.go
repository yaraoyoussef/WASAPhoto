/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

// errors
var ErrCouldNotModify = errors.New("could not modify username")
var ErrUserDoesNotExist = errors.New("could not find user")

// constante representing number of photos of each user in a home page of another user
const PPU = 2

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	Login(User) error
	GetProfile(User, User) (Profile, error)
	SetUsername(User, Username) error
	FollowUser(string, string) error
	UnfollowUser(string, string) error
	BanUser(string, string) error
	UnbanUser(string, string) error
	UploadPhoto(Photo) (int64, error)
	DeletePhoto(string, int64) error
	LikePhoto(int64, string) error
	UnlikePhoto(int64, string) error
	CommentPhoto(int64, string, Comment) (int64, error)
	UncommentPhotoOwner(int64, int64) error
	UncommentPhoto(string, int64, int64) error

	// utils methods
	GetUsername(string) (string, error)
	GetFollowers(string) ([]User, error)
	GetFollowing(string) ([]User, error)
	GetPhotos(string, string) ([]Photo, error)
	GetComments(string, string, int64) ([]Comment, error)
	GetLikes(string, string, int64) ([]User, error)
	CheckForBan(string, string) (bool, error)
	CheckForFollow(string, string) (bool, error)
	Ping() error
}

// structure that represents a user
type User struct {
	ID string
	// Username string
}

type Username struct {
	Username string
}

// structure that represents a user profile
type Profile struct {
	Username string `json:"username"`
	// photos are not strings and followers/following should be usernames not users
	Photos    []Photo `json:"photos"`
	Followers []User  `json:"followers"`
	Following []User  `json:"following"`
	Posts     int     `json:"posts"`
}

// structure that represents a photo uploaded
type Photo struct {
	ID          int       `json:"photoId"`
	Owner       string    `json:"owner"`
	Likes       []User    `json:"likes"`
	Comments    []Comment `json:"comments"`
	DateAndTime time.Time `json:"dateAndTime"`
}

// structure that represents a comment
type Comment struct {
	CommentId int64  `json:"commentId"`
	Comment   string `json:"comment"`
	UserId    string `json:"userId"`
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// enforce foreign keys (SQL doesn't do that by default)
	_, err := db.Exec(`PRAGMA foreign_keys = ON`)
	if err != nil {
		return nil, fmt.Errorf("error enforcing foreign keys: %w", err)
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	//var tableName string
	//err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='users';`).Scan(&tableName)
	//if errors.Is(err, sql.ErrNoRows) {
	err = createDB(db)
	if err != nil {
		return nil, fmt.Errorf("error creating database structure: %w", err)
	}
	//}

	return &appdbimpl{
		c: db,
	}, nil
}

// Create the database tables
func createDB(db *sql.DB) error {
	tables := [6]string{
		`CREATE TABLE IF NOT EXISTS users(
			id VARCHAR(16) NOT NULL PRIMARY KEY,
			username VARCHAR(16) NOT NULL);`,
		`CREATE TABLE IF NOT EXISTS banned(
			userId VARCHAR(16) NOT NULL,
			ubannedId VARCHAR(16) NOT NULL,
			PRIMARY KEY(userId, uBannedId),
			FOREIGN KEY(userId) REFERENCES users(id) ON DELETE CASCADE,
			FOREIGN KEY(uBannedId) REFERENCES users(id) ON DELETE CASCADE);`,
		`CREATE TABLE IF NOT EXISTS photos(
			photoId INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			userId VARCHAR(16) NOT NULL,
			dateAndTime DATETIME NOT NULL,
			FOREIGN KEY(userId) REFERENCES users(id) ON DELETE CASCADE);`,
		`CREATE TABLE IF NOT EXISTS followers(
			follower VARCHAR(16) NOT NULL, 
			followed VARCHAR(16) NOT NULL,
			PRIMARY KEY(follower, followed),
			FOREIGN KEY(follower) REFERENCES users(id) ON DELETE CASCADE,
			FOREIGN KEY(followed) REFERENCES users(id) ON DELETE CASCADE);`,
		`CREATE TABLE IF NOT EXISTS likes(
			photoId INTEGER NOT NULL,
			userId VARCHAR(16) NOT NULL, 
			PRIMARY KEY(photoId, userId),
			FOREIGN KEY(photoId) REFERENCES photos(photoId) ON DELETE CASCADE,
			FOREIGN KEY(userId) REFERENCES users(id) ON DELETE CASCADE);`,
		`CREATE TABLE IF NOT EXISTS comments(
			commentId INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			photoId INTEGER NOT NULL,
			userId VARCHAR(16) NOT NULL,
			comment VARCHAR(500) NOT NULL,
			FOREIGN KEY(photoId) REFERENCES photos(photoId) ON DELETE CASCADE,
			FOREIGN KEY(userId) REFERENCES users(id) ON DELETE CASCADE);`,
	}

	// executes all queries
	for i := 0; i < len(tables); i++ {
		_, err := db.Exec(tables[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
