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

type Args struct {
	Container     	string
	Kubernetes		bool
	Version			bool
	Logs			bool
	LogLength		string
}