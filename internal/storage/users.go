package storage

type UserFilters struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic"`
	Gender      string `json:"gender"`
	Age         string `json:"age"`
	Nationality string `json:"nationality"`
}

type User struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic,omitempty"`
	Gender      string `json:"gender"`
	Age         string `json:"age"`
	Nationality string `json:"nationality"`
}

type UsersCollection struct {
	Users []User `json:"users"`
}
