package user

import "database/sql"

type store struct {
	db *sql.DB
}

func CreateStore(db *sql.DB) *store {
	return &store{db: db}
}

func (s *store) CreateUser() {
}

func (s *store) GetUserByEmail() {

}
