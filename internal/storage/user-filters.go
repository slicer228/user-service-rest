package storage

import (
	"gorm.io/gorm"
)

type UserFilter struct {
	ID          []int
	Name        []string
	Surname     []string
	Patronymic  []string
	Gender      []string
	Age         []int
	Nationality []string
}

func filterUsers(s *gorm.DB, filter *UserFilter) *gorm.DB {
	if filter.ID != nil {
		s = s.Where("id IN ?", filter.ID)
	}
	if filter.Name != nil {
		s = s.Where("name IN ?", filter.Name)
	}
	if filter.Surname != nil {
		s = s.Where("surname IN ?", filter.Surname)
	}
	if filter.Patronymic != nil {
		s = s.Where("patronymic IN ?", filter.Patronymic)
	}
	if filter.Gender != nil {
		s = s.Where("gender IN ?", filter.Gender)
	}
	if filter.Age != nil {
		s = s.Where("age IN ?", filter.Age)
	}
	if filter.Nationality != nil {
		s = s.Where("nationality IN ?", filter.Nationality)
	}

	return s
}
