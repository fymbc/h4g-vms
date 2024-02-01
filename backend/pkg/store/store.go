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

func (s *Store) GetAllItems() ([]models.Item, error) {
    var items []models.Item
    rows, err := s.db.Query("SELECT id, name FROM items")
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
        var item models.Item
        if err := rows.Scan(&item.ID, &item.Name); err != nil {
            return nil, err
        }
        items = append(items, item)
    }
    return items, nil
}

func (s *Store) GetItem(id int) (*models.Item, error) {
    item := &models.Item{}
    err := s.db.QueryRow("SELECT id, name FROM items WHERE id = $1", id).Scan(&item.ID, &item.Name)
    if err != nil {
        return nil, err
    }
    return item, nil
}

func (s *Store) CreateItem(item *models.Item) error {
    err := s.db.QueryRow("INSERT INTO items (name) VALUES ($1) RETURNING id", item.Name).Scan(&item.ID)
    return err
}

func (s *Store) UpdateItem(item *models.Item) error {
    _, err := s.db.Exec("UPDATE items SET name = $1 WHERE id = $2", item.Name, item.ID)
    return err
}

func (s *Store) DeleteItem(id int) error {
    _, err := s.db.Exec("DELETE FROM items WHERE id = $1", id)
    return err
}
