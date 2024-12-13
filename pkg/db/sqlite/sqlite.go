package sqlite

import (
	"database/sql"
	"errors"
	"log/slog"
	"os"
	"todo/internal/config"
	"todo/pkg/logger"
	"todo/pkg/logger/sl"

	_ "github.com/mattn/go-sqlite3"
)

type sqliteDatabase struct {
	database *sql.DB
}

func (s *sqliteDatabase) Close() error {
	return s.database.Close()
}

func (s *sqliteDatabase) Prepare(query string) (*sql.Stmt, error) {
	return s.database.Prepare(query)
}

func (s *sqliteDatabase) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return s.database.Query(query, args...)
}

func Init(removeDB bool) *sqliteDatabase {
	path := config.GetConfig().Sqlite.Path

	if removeDB {
		logger.Info("Remove database " + path)
		if err := os.Remove(path); err != nil {
			logger.Error("can't remove database", sl.Err(err))
			return nil
		}
	}

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		logger.Warn("Database is not exists", slog.String("path", path))
		createDB(path)
	}

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		logger.Error("can't open"+path, sl.Err(err))
	}
	// TODO: ADD CLOSER
	err = db.Ping()
	if err != nil {
		logger.Fatal("Не могу достучаться до базы данных", sl.Err(err))
	}
	migrations(db)
	logger.Info("Connected to database", slog.String("path", path))
	return &sqliteDatabase{database: db}
}

func createDB(path string) {
	logger.Info("Creating database " + path)
	file, err := os.Create(path)
	if err != nil {
		logger.Fatal(path+" is not created", sl.Err(err))
	}
	err = file.Close()
	if err != nil {
		logger.Error("can't closing file", sl.Err(err))
	}
}

func migrations(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS tasks (
                id INTEGER PRIMARY KEY AUTOINCREMENT,
                title TEXT NOT NULL,
                description TEXT,
                status TEXT NOT NULL,
                created_at DATETIME DEFAULT CURRENT_TIMESTAMP)`

	_, err := db.Exec(query)
	if err != nil {
		logger.Fatal("can't create table tasks", sl.Err(err))
	}
	logger.Info("Table tasks created")
}
