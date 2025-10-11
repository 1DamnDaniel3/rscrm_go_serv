package dto

import "time"

type UserProfileCreateUpdateDTO struct {
	ID         string    `json:"id"`
	Account_id int64     `json:"account_id"`
	Phone      string    `json:"phone"`
	Full_name  string    `json:"full_name"`
	Birthdate  time.Time `json:"birthdate"`
}

type UserProfileResponseDTO struct {
	ID         string    `json:"id"`
	Account_id int64     `json:"account_id"`
	Phone      string    `json:"phone"`
	Full_name  string    `json:"full_name"`
	Birthdate  time.Time `json:"birthdate"`
}
