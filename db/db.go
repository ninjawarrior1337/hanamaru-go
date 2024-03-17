package db

import (
	"database/sql"
	"embed"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed migrations/*.sql
var migrations embed.FS

func InitDB(path string) (*sql.DB, error) {
	// if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
	// 	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	defer file.Close()
	// }
	// os.Create()
	return sql.Open("sqlite3", path)
}

func RunMigrations(db *sql.DB) error {
	fs, err := iofs.New(migrations, "migrations")
	if err != nil {
		return err
	}
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithInstance(
		"iofs",
		fs,
		"sqlite3",
		driver,
	)
	if err != nil {
		return err
	}
	err = m.Up()
	if err != nil {
		return err
	}
	fs.Close()
	return nil
}
