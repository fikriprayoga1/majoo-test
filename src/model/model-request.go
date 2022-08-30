package model

type ModelRequestLogin struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
