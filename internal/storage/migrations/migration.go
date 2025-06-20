package migrations

import "gorm.io/gorm"

func MustMigrate(gormDB *gorm.DB) {
	err := gormDB.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}
}
