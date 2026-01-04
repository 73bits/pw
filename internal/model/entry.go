package model

type Entry struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Notes    string `json:"notes,omitempty"`
}
