package models

type Register struct {
	Username string `form:"username"  binding:"required" json:"username,omitempty"`
	Password string `form:"password"  binding:"required" json:"password,omitempty"`
	Code     string `form:"code"  binding:"required" json:"code,omitempty"`
	Number   string `form:"number"  binding:"required" json:"number,omitempty"`
}
type Login struct {
	Username string `form:"username" json:"username,omitempty" binding:"required"`
	Password string `form:"password" json:"password,omitempty" binding:"required"`
	//Token string
}
type Forget struct {
	Username    string `json:"username,omitempty" form:"username"`
	NewPassword string `json:"newPassword,omitempty" form:"newPassword"`
	Code        string `json:"code,omitempty" form:"code"`
	Number      string `json:"number,omitempty" form:"number"`
}
