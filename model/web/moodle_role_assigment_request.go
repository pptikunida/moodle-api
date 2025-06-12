package web

type MoodleRoleAssigment struct {
	RoleID       int    `json:"roleid"`
	UserID       int    `json:"userid"`
	ContextID    int    `json:"contextid,omitempty"`
	ContextLevel string `json:"contextlevel,omitempty"`
	InstanceID   int    `json:"instanceid,omitempty"`
}

type MoodleRoleAssignRequest struct {
	Assignments []MoodleRoleAssigment `json:"assignments"`
}
