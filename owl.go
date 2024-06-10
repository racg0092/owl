package owl

import (
	"fmt"
	"github.com/racg0092/owl/pkg/models"
	"os"
	"path/filepath"
	"time"
)

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

func NewWatcher(path string, opt Options) (*Watcher, error) {
	stats, err := os.Stat(path)
	if err != nil {
		return &Watcher{}, err
	}
	abs, err := filepath.Abs(path)
	if err != nil {
		return &Watcher{}, err
	}

	w := &Watcher{
		Mode:    opt.Mode,
		Events:  make(chan Event),
		Errors:  make(chan error),
		Done:    make(chan int),
		Verbose: opt.Verbose,
	}

	var isFile bool
	if stats.IsDir() {

		w.Root = abs
		isFile = false
		w.Files = make([]models.File, 0)
		w.Directories = make([]string, 0)

	} else {

		w.File = models.File{
			Name:    stats.Name(),
			AbsLoc:  abs,
			ModTime: stats.ModTime(),
		}

		w.File.SetContext()
		isFile = true

	}

	if isFile {
		go fileProcess(w, w.File)
	} else {
		go dirProcess(w, w.Root)
	}

	return w, nil
}

func fileProcess(w *Watcher, file models.File) {
	for {
		if file.Context.Err() != nil {
			break
		}
		// some editors will rename or delete the old file when making changes therefore waiting half a second is recomended
		// todo: find a better sweet spot with the timeing
		time.Sleep(time.Millisecond * 500)
		if w.IsDone() {
			break
		}
		s, err := os.Stat(file.AbsLoc)
		if err != nil {
			w.Errors <- err
			continue
		}
		modTime := s.ModTime()
		if !file.ModTime.Equal(modTime) {

			if w.Verbose {
				fmt.Printf("[modified %v] %s\n", modTime.Format("15:04:05"), file.AbsLoc)
			}

			file.ModTime = modTime
			w.Events <- Event{
				Type:      FilEvent,
				Location:  file.AbsLoc,
				Operation: FileModified,
				Cancel:    file.Cancel,
			}

		}
	}
}

func dirProcess(w *Watcher, dir string) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		w.Errors <- err
		return
	}
	for _, entry := range entries {
		if entry.IsDir() {
			go dirRoutine(w, filepath.Join(dir, entry.Name()))
		} else {
			f, err := GetFileInfo(filepath.Join(dir, entry.Name()))
			if err != nil {
				w.Errors <- err
				continue
			}
			f.SetContext()
			go fileProcess(w, f)
		}
	}
}

func dirRoutine(w *Watcher, dir string) {
	mem := make(map[string]time.Time)
	for {
		time.Sleep(time.Millisecond * 500)
		entries, err := os.ReadDir(dir)
		if err != nil {
			w.Errors <- err
		}

		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}
			_, empty := mem[filepath.Join(dir, entry.Name())]
			if empty {
				continue
			}
			f, err := GetFileInfo(filepath.Join(dir, entry.Name()))
			if err != nil {
				w.Errors <- err
				continue
			}
			f.SetContext()
			go fileProcess(w, f)
			mem[filepath.Join(dir, entry.Name())] = f.ModTime
		}
	}
}
