package database

import (
	"database/sql"
	"davinci-chat/types"
	"log"
)

type Database interface {
	Close() error

	AddUser(request types.NewUserRequest) error
	ValidateUser(request types.ValidateUserRequest) error
}

func GetDatabase() Database {
	once.Do(func() {
		sqlite, err := sql.Open("sqlite3", "./user.sqlite")
		if err != nil {
			log.Fatalf("Failed to open database: %v", err)
		}

		sqliteDB = &SQLiteDB{db: sqlite}
		sqliteDB.createTables()
	})
	return sqliteDB
}
