package dto

import "time"

type StudentCreateUpdateDTO struct {
	Id          int64      `json:"id"`
	Name        string     `json:"name"`
	Birthdate   *time.Time `json:"birthdate"`
	Skill_level string     `json:"skill_level"`
	Contact     string     `json:"contact"`
	Created_at  time.Time  `json:"created_at"`
	School_id   string     `json:"school_id"`
}

type StudentResponseDTO struct {
	Id          int64      `json:"id"`
	Name        string     `json:"name"`
	Birthdate   *time.Time `json:"birthdate"`
	Skill_level string     `json:"skill_level"`
	Contact     string     `json:"contact"`
	Created_at  time.Time  `json:"created_at"`
	School_id   string     `json:"school_id"`
}
