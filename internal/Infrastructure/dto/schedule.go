package dto

import "time"

type ScheduleCreateUpdateDTO struct {
	ID               int64     `json:"id"`
	Group_id         int64     `json:"group_id"`
	Direction_id     int64     `json:"direction_id"`
	Teacher_id       int64     `json:"teacher_id"`
	Weekday          int       `json:"weekday"`
	Start_time       time.Time `json:"start_time"`
	Duration_minutes int       `json:"duration_minutes"`
	School_id        string    `json:"school_id"`
	Active_from      time.Time `json:"active_from"`
	Active_to        time.Time `json:"active_to"`
}

type ScheduleResponseDTO struct {
	ID               int64     `json:"id"`
	Group_id         int64     `json:"group_id"`
	Direction_id     int64     `json:"direction_id"`
	Teacher_id       int64     `json:"teacher_id"`
	Weekday          int       `json:"weekday"`
	Start_time       time.Time `json:"start_time"`
	Duration_minutes int       `json:"duration_minutes"`
	School_id        string    `json:"school_id"`
	Active_from      time.Time `json:"active_from"`
	Active_to        time.Time `json:"active_to"`
}
