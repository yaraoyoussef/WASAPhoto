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

var ErrCouldNotModify = errors.New("could not modify username")
var ErrUserDoesNotExist = errors.New("could not find user")

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	Login(User) error
	GetProfile(string) (Profile, error)
	SetUsername(User) error
	FollowUser(string, string) error
	UnfollowUser(string, string) error
	GetName() (string, error)
	SetName(string) error
	BanUser(string, string) error
	UnbanUser(string, string) error
	UploadPhoto(Photo) (int64, error)
	CheckForBan(string, string) (bool, error)
	DeletePhoto(string, int64) error
	LikePhoto(int64, string) error
	UnlikePhoto(int64, string) error
	CommentPhoto(int64, string, Comment) (int64, error)
	UncommentPhotoOwner(int64, int64) error
	UncommentPhoto(string, int64, int64) error
	Ping() error
}

// structure that represents a user
type User struct {
	ID       string
	Username string
	// each user has to a profile
	//Pr Profile
}

// structure that represents a user profile
type Profile struct {
	Username string `json:"username"`
	// photos are not strings and followers/following should be usernames not users
	Photos    []Photo  `json:"photos"`
	Followers []string `json:"followers"`
	Following []string `json:"following"`
	Posts     int      `json:"posts"`
}

type Photo struct {
	ID          int       `json:"photoId"`
	Owner       string    `json:"owner"`
	Likes       []User    `json:"likes"`
	Comments    []Comment `json:"comments"`
	DateAndTime time.Time `json:"dateAndTime"`
}

type Comment struct {
	CommentId int64  `json:"commentId"`
	Comment   string `json:"comment"`
	Username  string `json:"username"`
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

	// create tables here

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='example_table';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE example_table (id INTEGER NOT NULL PRIMARY KEY, name TEXT);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}

// create function to authenticate current user

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
