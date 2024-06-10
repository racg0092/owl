package tests

import (
	"fmt"
	"testing"

	"github.com/racg0092/owl"
)

func TestFileWatch(t *testing.T) {
	w, err := owl.NewWatcher("../sandbox/file1")
	if err != nil {
		t.Error(err)
	}
	for {
		select {
		case e, open := <-w.Events:
			if !open {
				return
			}
			if e == owl.FileModified {
				fmt.Println("File was modified")
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
