package user_manager

import "user-service/internal/storage"

type IUserManager interface {
	GetUsers(filter *storage.UserFilter, pag *storage.Paginate) []storage.User
	DeleteUsers(filter *storage.UserFilter) error
	AddUser(data *storage.UserData) (int, error)
	PatchUsers(data *storage.UserData, filter *storage.UserFilter) error
}
