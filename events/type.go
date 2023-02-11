package events

type Fetcher interface {
	Fetch(limit int) ([]Event, error)
}

type Processor interface {
	Process(e Event) error
}

type EventType int

const (
	Unknown EventType = iota
	Message
)

type Event struct {
	Type EventType
	Text string
	Meta interface{}
}
