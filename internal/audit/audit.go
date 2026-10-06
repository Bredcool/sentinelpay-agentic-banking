package audit

import "time"

type Event struct {
	ID         string    `json:"id"`
	IncidentID string    `json:"incident_id"`
	Type       string    `json:"type"`
	Message    string    `json:"message"`
	CreatedAt  time.Time `json:"created_at"`
}

type Logger interface {
	Record(event Event) error
}

type MemoryLogger struct {
	Events []Event
}

func (l *MemoryLogger) Record(event Event) error {
	l.Events = append(l.Events, event)
	return nil
}
