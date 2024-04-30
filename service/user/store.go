package user

import (
	"database/sql"
	"fmt"

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
	u := new(types.User)
	for rows.Next() {
		u, err = ScanRowIntoUser(rows)
		if err != nil {
			return nil,err
		}
	}
	if u.ID == 0 { // no such user in the database	
		return nil, fmt.Errorf("No user with this email found")
	}
		
	
	
	return u ,nil
}



func ScanRowIntoUser(row *sql.Rows) (*types.User, error) {
	user := new(types.User)
	err := row.Scan(&user.ID,
		   &user.FirstName,
		   &user.LastName,
		   &user.Password,
		   &user.Email)
	if  err != nil {
		return nil, err
	}
	return user, nil
}

func  (s *Store) CreateUser( user types.User) error {
	_, err := s.db.Exec("INSERT INTO users(firstname,lastname,email,password) VALUES(?,?,?,?) ",
			user.FirstName, user.LastName, user.Email, user.Password )
	if err != nil {
		return err
	}
	return  nil
}