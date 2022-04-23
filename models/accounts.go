package models

import "time"

type Account struct {
	ID       uint64    `json:"id"`
	Name     string    `json:"name"`
	DOB      time.Time `json:"dob"`
	EMail    string    `json:"email"`
	Contact  string    `json:"contact"`
	Password string    `json:"pwd"`
}
