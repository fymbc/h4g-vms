package store

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Store struct {
    Db *gorm.DB
}



func Connect() (*gorm.DB, error) {
    dsn := "dbname=postgres user=postgres password=postgres host=localhost port=5432 sslmode=disable"
    dbDriver := postgres.Open(dsn)

    db, err := gorm.Open(dbDriver, &gorm.Config{})
    if err != nil {
        return nil, err
    }
    return db, nil
}

func GetDB() (*gorm.DB, error) {
    db, err := Connect()
    if err != nil {
        return nil, err
    }
    return db, nil
}



// func (s *Store) GetAllItems() ([]models.Item, error) {
//     var items []models.Item
//     rows, err := s.db.Query("SELECT id, name FROM items")
//     if err != nil {
//         return nil, err
//     }
//     defer rows.Close()
//     for rows.Next() {
//         var item models.Item
//         if err := rows.Scan(&item.ID, &item.Name); err != nil {
//             return nil, err
//         }
//         items = append(items, item)
//     }
//     return items, nil
// }

// func (s *Store) GetItem(id int) (*models.Item, error) {
//     item := &models.Item{}
//     err := s.db.QueryRow("SELECT id, name FROM items WHERE id = $1", id).Scan(&item.ID, &item.Name)
//     if err != nil {
//         return nil, err
//     }
//     return item, nil
// }

// func (s *Store) CreateItem(item *models.Item) error {
//     err := s.db.QueryRow("INSERT INTO items (name) VALUES ($1) RETURNING id", item.Name).Scan(&item.ID)
//     return err
// }

// func (s *Store) UpdateItem(item *models.Item) error {
//     _, err := s.db.Exec("UPDATE items SET name = $1 WHERE id = $2", item.Name, item.ID)
//     return err
// }

// func (s *Store) DeleteItem(id int) error {
//     _, err := s.db.Exec("DELETE FROM items WHERE id = $1", id)
//     return err
// }


// func (s *Store) AuthenticateUser(username, password string) (*models.User, error) {
//     // This is a simplified example. In a real application, passwords should be hashed.
//     // NEVER store or compare plain text passwords!
//     var user models.User
//     err := s.db.QueryRow("SELECT id, username, password FROM users WHERE username = $1", username).Scan(&user.ID, &user.Username, &user.Password)
//     if err != nil {
//         return nil, err // User not found or other error
//     }

//     // Here, compare the hashed password from the DB with the hashed version of the input password.
//     // if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
//     //     return nil, err // Password does not match
//     // }

//     return &user, nil
// }
