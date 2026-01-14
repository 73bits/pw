package model

type Entry struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Vault struct {
	Entries []Entry `json:"entries"`
}
