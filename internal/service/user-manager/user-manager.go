package user_manager

import (
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"log/slog"
	"user-service/internal/config"
	age "user-service/internal/lib/api/supposing-api/suppose-age"
	gender "user-service/internal/lib/api/supposing-api/suppose-gender"
	nation "user-service/internal/lib/api/supposing-api/suppose-nationality"
	"user-service/internal/storage"
)

type UserManager struct {
	IUserManager
	db *storage.DBConnection
}

type PrimaryUserData struct {
	Name       string `validate:"required"`
	Surname    string `validate:"required"`
	Patronymic string
}

func (um *UserManager) GetUsers(log *slog.Logger, filter *storage.UserFilter, pag *storage.Paginate) []storage.User {
	s := um.db.NewSession(log)
	return s.GetUsers(filter, pag)
}

func (um *UserManager) DeleteUsers(log *slog.Logger, filter *storage.UserFilter) error {
	s := um.db.NewSession(log)
	err := s.DeleteUsers(filter)
	if err != nil {
		return err
	}
	return nil
}

func (um *UserManager) AddUser(log *slog.Logger, data *PrimaryUserData) (int, error) {
	s := um.db.NewSession(log)

	var userData *storage.UserData

	copier.Copy(&userData, &data)

	log.Debug("Fetching supposed age...")
	userData.Age = age.RequestPredictedAge(log, userData.Name)
	log.Debug("Fetching supposed nationality...")
	userData.Nationality = nation.RequestPredictedNationality(log, userData.Name)
	log.Debug("Fetching supposed gender...")
	userData.Gender = gender.RequestPredictedGender(log, userData.Name)

	userId, err := s.CreateUser(userData)
	if err != nil {
		return 0, err
	}
	return userId, nil
}

func (um *UserManager) PatchUsers(log *slog.Logger, data *storage.UserData, filter *storage.UserFilter) error {
	s := um.db.NewSession(log)
	err := s.PatchUsers(data, filter)

	if err != nil {
		return err
	}
	return nil
}

func NewUserManager(db *gorm.DB, log *slog.Logger, cfg *config.Config) *UserManager {
	return &UserManager{
		db: storage.MustLoadDB(cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPassword, cfg.DbName),
	}
}
