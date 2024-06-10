package owl

const (
	FilEvent = iota
	DirEvent
)

type Event struct {

	// Defines if is realted to a file of a directory
	// I think this can be remved and just use operation
	Type int

	// Absolute location
	Location string

	// Type of operation performed on the file
	Operation int

	// Cancels the go routine for the current event
	Cancel func()
}
