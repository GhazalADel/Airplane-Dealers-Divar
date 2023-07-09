package datastore

import "Airplane-Divar/models"

type User interface {
	Get(id int) ([]models.User, error)
	Create(username string, is_admin bool, password string) (string, models.User, error)
	Login(username, password string, is_admin bool) (string, models.User, error)
	CheckUnique(username string) (string, error)
}

type Payment interface {
	Create(userID uint, fee int64, authority string) (string, error)
}
