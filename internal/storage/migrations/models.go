package migrations

type User struct {
	ID          int `gorm:"primaryKey"`
	Name        string
	Surname     string
	Patronymic  *string
	Age         int
	Gender      string
	Nationality string
}
