package storage

import (
	"github.com/jinzhu/copier"
	"user-service/internal/storage/migrations"
)

type UserData struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic"`
	Gender      string `json:"gender"`
	Age         int    `validate:"gte=0,lte=130" json:"age"`
	Nationality string `json:"nationality"`
}

type User struct {
	UserData
	ID int `json:"id"`
}

type Paginate struct {
	Page         int
	ItemsPerPage int
}

// Returns user_id if succeed
func (s *DBSession) CreateUser(user *UserData) (int, error) {
	s.log.Debug("Storage creating user...")
	if err := s.v.Struct(user); err != nil {
		s.log.Error("Validation data error", "err", err.Error())
		return 0, err
	}

	var dbUser migrations.User
	copier.Copy(&dbUser, &user)

	s.session.Create(&dbUser)

	return dbUser.ID, nil
}

func (s *DBSession) DeleteUser(userIds []int) error {
	s.log.Debug("Storage deleting users...")
	var dbUser migrations.User

	filterUsers(s.session, &UserFilter{ID: userIds}).Delete(&dbUser)

	return nil
}

func (s *DBSession) PatchUser(updateParams *UserData, userIds []int) error {
	s.log.Debug("Storage patching users...")
	if err := s.v.Struct(updateParams); err != nil {
		s.log.Error("Validation data error", "err", err.Error())
		return err
	}

	var dbUser migrations.User
	copier.Copy(&dbUser, &updateParams)

	filterUsers(s.session.Model(&migrations.User{}), &UserFilter{ID: userIds}).Updates(dbUser)

	return nil
}

func (s *DBSession) GetUsers(filter *UserFilter, pag *Paginate) []User {
	s.log.Debug("Storage fetching users...")
	var dbUsers []migrations.User

	paginate(filterUsers(s.session, filter), pag.Page, pag.ItemsPerPage).Find(&dbUsers)

	users := make([]User, len(dbUsers))

	for i, dbUser := range dbUsers {
		copier.Copy(&users[i], dbUser)
	}

	return users
}
