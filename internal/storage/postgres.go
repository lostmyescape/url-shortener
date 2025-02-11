package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/user/urlProject/config"
	"log"
)

type Storage struct {
	DB *sql.DB
}

// NewStorage соберет и вернет объект storage
func NewStorage(cfg *config.Config) (*Storage, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Storage.Host,
		cfg.Storage.Port,
		cfg.Storage.User,
		cfg.Storage.Password,
		cfg.Storage.DbName,
		cfg.Storage.SslMode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения к БД: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("не удалось подключиться к БД: %w", err)
	}
	log.Println("Подключение к postgres установлено!")

	// Выполняем SQL-запрос на создание таблицы
	createTable := `
    CREATE TABLE IF NOT EXISTS url (
        id SERIAL PRIMARY KEY,
        alias TEXT NOT NULL UNIQUE,
        url TEXT NOT NULL
    );
    CREATE INDEX IF NOT EXISTS idx_alias ON url(alias);
    `
	_, err = db.Exec(createTable)
	if err != nil {
		return nil, fmt.Errorf("ошибка при создании таблицы url: %w", err)
	}

	return &Storage{DB: db}, nil
}

func (s *Storage) SaveUrl(urlToSave string, alias string) (int64, error) {
	const op = "storage.sqlite.SaveUrl"

	var id int64

	query := "INSERT INTO url(url, alias) VALUES ($1, $2) RETURNING id"

	err := s.DB.QueryRow(query, urlToSave, alias).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil

}
