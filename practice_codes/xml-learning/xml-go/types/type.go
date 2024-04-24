package types

type User struct {
	UserName string `json:"usename,omitempty"`
	Password string `json:"usename,omitempty"`
	Email    string `json:"usename,omitempty"`
}

type UserDb struct {
	Users []User `json:"users,omitempty"`
	Type  string `json:"type,omitempty"`
}
