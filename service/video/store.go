package video

import "database/sql"

type store struct {
	db *sql.DB
}

func CreateStore(db *sql.DB) *store {
	return &store{db: db}
}

func (s *store) UploadtoS3() string {
	return "bruh"
}
