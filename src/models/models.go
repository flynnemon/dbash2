package models

import (
	"time"
)
type Container struct {
	Name     	string
	CreatedAt 	time.Time
	Image  		string
	ID			string
	State		string
}
