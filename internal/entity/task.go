package entity

type Status string

const (
	StatusQueued  = "queued"
	StatusRunning = "running"
	StatusDone    = "done"
	StatusFailed  = "failed"
)

type Task struct {
	ID         string
	Payload    string
	MaxRetries int
	Attempts   int32
}
