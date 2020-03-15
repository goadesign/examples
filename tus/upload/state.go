package upload

// State indicates the state of an upload.
type State int

const (
	Started State = iota
	Completed
	Cancelled
	TimedOut
	Failed
)

var names = [...]string{
	"started",
	"completed",
	"cancelled",
	"timed out",
	"failed",
}

// String returns a human-friendly description of the upload state.
func (s State) String() string {
	return names[s]
}
