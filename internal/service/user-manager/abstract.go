package user_manager

type IUserManager interface {
	GetUsers()
	DeleteUser()
	AddUser()
	PatchUser()
}
