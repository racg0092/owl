package models

import (
	"time"
)

type File struct {
	Name string
	// Absolute location
	AbsLoc  string
	ModTime time.Time
}
