package owl

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/racg0092/owl/pkg/models"
)

var outputChannel = make(chan int)

// Gets File Information. Takes in a `loc` (absolute or relative) and returns
// the file information or an error
func GetFileInfo(loc string) (models.File, error) {
	stats, err := os.Stat(loc)
	if err != nil {
		return models.File{}, err
	}
	abs, err := filepath.Abs(loc)
	if err != nil {
		return models.File{}, err
	}
	if stats.IsDir() {
		return models.File{}, fmt.Errorf("Resource is a directory not a file")
	}
	return models.File{Name: stats.Name(), ModTime: stats.ModTime(), AbsLoc: abs}, nil
}

func NewWatcher(path string) (*Watcher, error) {
	stats, err := os.Stat(path)
	if err != nil {
		return &Watcher{}, err
	}
	abs, err := filepath.Abs(path)
	if err != nil {
		return &Watcher{}, err
	}

	w := &Watcher{
		Mode:   PollingMode,
		Events: make(chan int),
		Errors: make(chan error),
		Done:   make(chan int),
	}
	var isFile bool
	if stats.IsDir() {
		w.Root = abs
		isFile = false
	} else {
		w.File = models.File{
			Name:    stats.Name(),
			AbsLoc:  abs,
			ModTime: stats.ModTime(),
		}
		isFile = true
	}

	if isFile {
		go fileProcess(w)
	} else {
		go dirProcess(w)
	}

	return w, nil
}

func fileProcess(w *Watcher) {
	for {
		// some editors will rename or delete the old file when making changes therefore waiting half a second is recomended
		// todo: find a better sweet spot with the timeing
		time.Sleep(time.Millisecond * 500)
		if w.IsDone() {
			break
		}
		s, err := os.Stat(w.File.AbsLoc)
		if err != nil {
			w.Errors <- err
			continue
		}
		modTime := s.ModTime()
		if !w.File.ModTime.Equal(modTime) {
			w.File.ModTime = modTime
			w.Events <- FileModified
		}
	}
}

func dirProcess(w *Watcher) {

}
