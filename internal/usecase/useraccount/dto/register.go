package dto

import "time"

// RegisterInput входные данные для регистрации
type RegisterInput struct {
	School  SchoolDTO  `json:"school" example:"{\"name\":\"Right Step\",\"city\":\"Тимашевск\",\"phone\":\"+7-999-123-45-67\",\"email\":\"popov@gmail.com\"}"`
	Account AccountDTO `json:"account" example:"{\"email\":\"popov@gmail.com\",\"password\":\"secret\"}"`
	Profile ProfileDTO `json:"profile" example:"{\"full_name\":\"Артём Попов\",\"phone\":\"+7-999-123-45-67\",\"birthdate\":\"1985-01-01T00:00:00Z\"}"`
}

// RegisterOutput данные, возвращаемые после успешной регистрации
type RegisterOutput struct {
	School string `json:"school" example:"Right Step"`
}

// SchoolDTO данные о школе
type SchoolDTO struct {
	Name  string `json:"name" example:"Right Step"`
	City  string `json:"city" example:"Тимашевск"`
	Phone string `json:"phone" example:"+7-999-123-45-67"`
	Email string `json:"email" example:"popov@gmail.com"`
}

// AccountDTO данные для аккаунта
type AccountDTO struct {
	Email    string `json:"email" example:"popov@gmail.com"`
	Password string `json:"password" example:"secret"`
}

// ProfileDTO данные профиля пользователя
type ProfileDTO struct {
	FullName  string    `json:"full_name" example:"Артём Попов"`
	Phone     string    `json:"phone" example:"+7-999-123-45-67"`
	Birthdate time.Time `json:"birthdate" example:"1985-01-01T00:00:00Z"`
}
