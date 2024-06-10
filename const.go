package owl

const (
	// Watches the directory or file through short polling method
	PollingMode = iota

	// Watches the directory of firl throuhg system events
	SystemEventsMode
)

const (
	FileModified = iota
	FileCreated
	DirectoryAdded
	FileRemoved
)
