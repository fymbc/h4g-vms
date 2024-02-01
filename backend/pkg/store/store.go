package store

import (
	"database/sql"
	"h4g-vms/pkg/models"
)

type Store struct {
    db *sql.DB
}

func New(databaseURL string) (*Store, error) {
    db, err := sql.Open("postgres", databaseURL)
    if err != nil {
        return nil, err
    }
    return &Store{db: db}, nil
}

func (s *Store) GetItems() ([]models.Item, error) {
    // Implement database access to retrieve items
    // ...
    return nil, nil
}
