package model

import (
	"time"
)

type Trip struct {
	Id          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Start_date  time.Time `json:"start_date"`
	End_date    time.Time `json:"end_date"`
	CreatedAt   time.Time `json:"created_at"`
	IsPublic    bool      `json:"inPublic"`
	// User        User      `json:"user"`
	User_Id     uint      `json:"user_id"`
}
