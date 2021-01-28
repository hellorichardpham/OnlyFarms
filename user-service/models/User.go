package models

import "time"

//User MySQL db table
type User struct {
	ID       int
	Name     string
	Email    string
	Password string
	Created  time.Time
	Active   bool
}
