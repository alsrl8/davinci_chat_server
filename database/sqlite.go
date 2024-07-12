package database

import (
	"database/sql"
	"davinci-chat/err_types"
	"davinci-chat/types"
	"errors"
	"log"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteDB struct {
	db *sql.DB
}

var (
	sqliteDB *SQLiteDB
	once     sync.Once
)

func (s *SQLiteDB) Close() error {
	return s.db.Close()
}

func (s *SQLiteDB) AddUser(request types.NewUserRequest) error {
	return s.createUser(request)
}

func (s *SQLiteDB) ValidateNewUser(request types.ValidateUserRequest) error {
	return s.validateUser(request)
}

func (s *SQLiteDB) Login(request types.LoginRequest) (string, error) {
	return s.login(request)
}

func (s *SQLiteDB) createTables() {
	createUserTable := `CREATE TABLE IF NOT EXISTS user (
       "email" TEXT NOT NULL PRIMARY KEY,
       "name" TEXT NOT NULL,
       "password" TEXT NOT NULL
   );`
	_, err := s.db.Exec(createUserTable)
	if err != nil {
		log.Fatalf("Error creating user table: %v", err)
	}
}

func (s *SQLiteDB) createUser(request types.NewUserRequest) error {
	createUser := `INSERT INTO user (name, email, password) VALUES (?, ?, ?);`
	_, err := s.db.Exec(createUser, request.UserName, request.UserEmail, request.Password)
	return err
}

func (s *SQLiteDB) validateUser(request types.ValidateUserRequest) error {
	getUserByEmail := `SELECT EXISTS(SELECT 1 FROM user WHERE email = ?);`
	row := s.db.QueryRow(getUserByEmail, request.UserEmail)

	var exists bool
	if err := row.Scan(&exists); err != nil {
		return err
	}

	if exists {
		return err_types.ErrUserExists
	}

	return nil
}

func (s *SQLiteDB) login(request types.LoginRequest) (string, error) {
	email := request.UserEmail
	password := request.Password

	loginQuery := `SELECT name FROM user WHERE email = ? and password = ?;`
	row := s.db.QueryRow(loginQuery, email, password)

	var name string
	if err := row.Scan(&name); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", err_types.ErrEmailPasswordNotMatch
		}
		return "", err
	}

	return name, nil
}
