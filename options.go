package owl

// Options for initializing the watcher
type Options struct {
	// If verbose is set to true then most of the events get printed to the stdout
	Verbose bool

	// Which mode you want the watcher to be on. [PollingMode] or [SystemEventsMode]
	Mode int
}
