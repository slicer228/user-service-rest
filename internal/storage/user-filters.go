package storage

import (
	"gorm.io/gorm"
)

type UserFilter struct {
	ID          []int    `form:"id"`
	Name        []string `form:"name"`
	Surname     []string `form:"surname"`
	Patronymic  []string `form:"patronymic"`
	Gender      []string `form:"gender"`
	Age         []int    `form:"age"`
	Nationality []string `form:"nationality"`
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
