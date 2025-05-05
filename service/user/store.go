package user

import (
	"database/sql"
	"fmt"

	"github.com/xadhithiyan/videon/types"
)

type store struct {
	db *sql.DB
}

func CreateStore(db *sql.DB) *store {
	return &store{db: db}
}

func (s *store) CreateUser(user types.User) error {
	_, err := s.db.Query(`INSERT INTO "user" (name, email, password) VALUES($1, $2, $3)`,
		user.Name, user.Email, user.Password)

	if err != nil {
		return err
	}
	return nil
}

func (s *store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query(`SELECT * FROM "user" WHERE email = $1`, email)
	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		err := rows.Scan(
			&u.Id,
			&u.Name,
			&u.Email,
			&u.Password,
		)

		if err != nil {
			return nil, err
		}
	}

	if u.Id == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}
