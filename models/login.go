package models

type Admin struct {
	Name     string `json:"username,omitempty" form:"username"`
	Password string `json:"password,omitempty" form:"password"`
}
