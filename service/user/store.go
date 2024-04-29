package user

import (
	"database/sql"

	"github.com/Magdi888/EcoWithGo/types"
)

type Store struct {
	db *sql.DB
}

func  NewStore(db *sql.DB) *Store{
	return &Store{db: db}
}

// GetUserByEmail returns a User
func (s *Store) GetUserByEmail (email string) (*types.User, error){
	rows, err := s.db.Query("SELECT email FROM  users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		ScanRowIntoUser(rows)
	}
}


func ScanRowIntoUser(row *sql.Rows) (*types.User, error) {
	var user  types.User
	err := row.Scan(&user.ID,
		   &user.FirstName,
		   &user.LastName,
		   &user.Password,
		   &user.Email)
	return &user, err
}