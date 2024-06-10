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

	// Context passed to the file routine
	Context context.Context

	// Cancels the context for the file watcher routine. It only affects the
	// file calling it. All other routines stay active.
	Cancel func()
}

// Initializes a context with cancel function
func (f *File) SetContext() {
	f.Context, f.Cancel = context.WithCancel(context.Background())
}
