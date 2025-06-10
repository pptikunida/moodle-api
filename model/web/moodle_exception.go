package web

import "fmt"

type MoodleException struct {
	Exception string `json:"exception"`
	ErrorCode string `json:"errorcode"`
	Message   string `json:"message"`
	DebugInfo string `json:"debuginfo,omitempty"`
}

func (e *MoodleException) Error() string {
	return fmt.Sprintf("moodle API error: %s", e.Exception)
}
