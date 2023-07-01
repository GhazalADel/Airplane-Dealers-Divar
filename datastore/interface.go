package datastore

import "Airplane-Divar/models"

type User interface {
	Get(id int) ([]models.User, error)
	Create(user models.User) (models.User, error)
	CheckUnique(user models.User) (string, error)
}

type Account interface {
	Get(id int) ([]models.Account, error)
	Create(user_id int, username string, is_admin bool, password string) (string, models.Account, error)
	Login(username, password string, is_admin bool) (string, models.Account, error)
	CheckUnique(username string) (string, error)
}
