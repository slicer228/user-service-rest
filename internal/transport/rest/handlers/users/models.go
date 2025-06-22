package users

import (
	"user-service/internal/service/user-manager"
	"user-service/internal/storage"
)

type RequestAddUser struct {
	user_manager.PrimaryUserData
}

type RequestPatchUser struct {
	storage.UserData
	UserIds []int `json:"user_ids"`
}

type RequestDeleteUser struct {
	UserIds []int `json:"user_ids"`
}

type RequestGetUsers struct {
	storage.UserFilter
	storage.Paginate
}
