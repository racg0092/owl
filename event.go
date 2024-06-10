package owl

const (
	FilEvent = iota
	DirEvent
)

type Event struct {
	Type      int
	Location  string
	Operation int
	Cancel    func()
}
