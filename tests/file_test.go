package tests

import (
	"fmt"
	"testing"

	"github.com/racg0092/owl"
)

func TestFileWatch(t *testing.T) {
	w, err := owl.NewWatcher("../sandbox/file1", owl.Options{})
	if err != nil {
		t.Error(err)
	}
	for {
		select {
		case e, open := <-w.Events:
			if !open {
				return
			}
			if e.Operation == owl.FileModified {
				fmt.Printf("File was modified. %s\n", e.Location)
			} else {
				fmt.Printf("%v Something else happend to the file\n", e)
			}
		case err, open := <-w.Errors:
			if !open {
				return
			}
			t.Error(err)
		case _ = <-w.Done:
			fmt.Println("Process is done")
		}
	}
}
