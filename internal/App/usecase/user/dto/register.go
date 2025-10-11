package dto

import "time"

// RegisterInput входные данные для регистрации
type RegisterInput struct {
	School  SchoolDTO  `json:"school"`
	Account AccountDTO `json:"account"`
	Profile ProfileDTO `json:"profile"`
}

// RegisterOutput данные, возвращаемые после успешной регистрации
type RegisterOutput struct {
	School_id string `json:"school_id" example:"bbeb26e7-7a3a-4bcf-8a70-338f362eabd1"`
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
