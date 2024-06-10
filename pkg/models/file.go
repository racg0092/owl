package models

import (
	"context"
	"time"
)

type File struct {
	Name string

	// Absolute location
	AbsLoc string

	ModTime time.Time

	Context context.Context

	Cancel func()
}

func (f *File) SetContext() {
	f.Context, f.Cancel = context.WithCancel(context.Background())
}
