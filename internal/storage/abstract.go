package storage

import "log/slog"

type Connection interface {
	NewSession(log *slog.Logger) *DBSession
}

type Session interface {
	CreateUser(user UserData) (int, error)
	DeleteUsers(filter UserFilter) error
	PatchUsers(updateParams UserData, filter UserFilter) error
	GetUsers(filter UserFilter, pag Paginate) []user
}
