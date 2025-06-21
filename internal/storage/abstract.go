package storage

import "log/slog"

type Connection interface {
	NewSession(log *slog.Logger) *DBSession
}

type Session interface {
	CreateUser(user UserData) (int, error)
	DeleteUser(userId int) error
	PatchUser(updateParams *UserData, userId int) error
	GetUsers(filter UserFilter, pag Paginate) []User
}
