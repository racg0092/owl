package owl

import (
	"time"

	"github.com/racg0092/owl/pkg/models"
)

type Watcher struct {
	done bool

	// How often should the files be checked. Only applies when [Mode] is polling
	ShortPollingInterval time.Time

	// Mode of wacther
	Mode int

	// If your watching a single file this is the absolute path. Otherwise the value of [File] is ""
	File models.File

	// If your watching a directory this is the root directitory
	Root string

	// List of files being watched
	Files []models.File

	// Event channels
	Events chan int

	// Error channels
	Errors chan error

	// Directories in the [Root] is any
	Directories []string

	Done chan int
}

func (w *Watcher) Close() {
	w.done = true
	w.Done <- 1
	close(w.Events)
	close(w.Errors)
}

func (w *Watcher) IsDone() bool {
	return w.done
}
