package web

type MoodleUserGetByFieldRequest struct {
	Field  string   `json:"field"`  // "id", "idnumber", "username", "email"
	Values []string `json:"values"` // Example: ["john.doe"]
}
