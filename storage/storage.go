package storage

import (
	"github.com/nurmuhammaddeveloper/mudium_user_service/storage/postgres"

	"github.com/jmoiron/sqlx"
	"github.com/nurmuhammaddeveloper/mudium_user_service/storage/repo"
)

type StorageI interface {
	User() repo.UserStorageI
}

type storagePg struct {
	userRepo repo.UserStorageI
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
		userRepo: postgres.NewUser(db),
	}
}

func (s *storagePg) User() repo.UserStorageI {
	return s.userRepo
}
