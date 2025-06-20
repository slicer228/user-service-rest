package storage

import (
	"github.com/jinzhu/copier"
	"user-service/internal/storage/migrations"
)

type UserData struct {
	Name        string
	Surname     string
	Patronymic  string
	Gender      string
	Age         int `validate:"gte=0,lte=130"`
	Nationality string
}

type User struct {
	UserData
	ID int
}

type Paginate struct {
	Page         int
	ItemsPerPage int
}

// Returns user_id if succeed
func (s *DBSession) CreateUser(user *UserData) (int, error) {
	if err := s.v.Struct(user); err != nil {
		s.log.Error("Validation data error", "err", err.Error())
		return 0, err
	}

	var dbUser migrations.User
	copier.Copy(&dbUser, &user)

	s.session.Create(&dbUser)

	return dbUser.ID, nil
}

func (s *DBSession) DeleteUsers(filter *UserFilter) error {

	var dbUser migrations.User

	filterUsers(s.session, filter).Delete(&dbUser)

	return nil
}

func (s *DBSession) PatchUsers(updateParams *UserData, filter *UserFilter) error {
	if err := s.v.Struct(updateParams); err != nil {
		s.log.Error("Validation data error", "err", err.Error())
		return err
	}

	var dbUser migrations.User
	copier.Copy(&dbUser, &updateParams)

	filterUsers(s.session.Model(&migrations.User{}), filter).Updates(dbUser)

	return nil
}

func (s *DBSession) GetUsers(filter *UserFilter, pag *Paginate) []User {

	var dbUsers []migrations.User

	paginate(filterUsers(s.session, filter), pag.Page, pag.ItemsPerPage).Find(&dbUsers)

	users := make([]User, len(dbUsers))

	for i, dbUser := range dbUsers {
		copier.Copy(&users[i], dbUser)
	}

	return users
}
