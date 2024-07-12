package database

import (
	"database/sql"
	"davinci-chat/config"
	"davinci-chat/types"
	"log"
)

type Database interface {
	Close() error

	AddUser(request types.NewUserRequest) error
	ValidateUser(request types.ValidateUserRequest) error
}

func GetUserDatabase() Database {
	once.Do(func() {
		dsName := config.GetDataSourceName("user")
		sqlite, err := sql.Open("sqlite3", dsName)
		if err != nil {
			log.Fatalf("Failed to open database: %v", err)
		}

		sqliteDB = &SQLiteDB{db: sqlite}
		sqliteDB.createTables()
	})
	return sqliteDB
}
