package activities

import (
	"time"
)

type Environment int

const (
	Development Environment = iota
	Testing
	Staging
	Production
)

type Activity struct {
	Name          string // function name or url path
	Authorization string
	UserID        UserID
	Login         string
	SessionID     string
	TraceID       string
	Client        Client
	Environment   Environment
	StartTime     time.Time
}
