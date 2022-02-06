package models

type Admin struct {
	Name     string `json:"name,omitempty" form:"name"`
	Password string `json:"password,omitempty" form:"password"`
}
