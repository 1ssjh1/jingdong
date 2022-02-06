package models

type BasicInfo struct {
	Username string
	Uid      int
}
type Balance struct {
	Username string `form:"username" json:"username,omitempty" `
	Balance  string `form:"balance" json:"balance,omitempty" `
	Token    string `form:"token" json:"token" `
}
type User struct {
	Username string `json:"username,omitempty" form:"username"`
	Token    string `json:"token,omitempty" form:"token"`
}
type UserOrder struct {
	Uid      int
	Allorder []AllOrder
}
type AllOrder struct {
	Gid   int
	Oid   int
	Cid   int
	Count int
	State string
}
type Commit struct {
	Uid      int    `json:"uid,omitempty" form:"uid"`
	Username string `json:"username,omitempty" form:"username"`
	Token    string `json:"token,omitempty" form:"token"`
	Oid      int    `json:"oid,omitempty" form:"oid"`
	Commit   string `json:"commit,omitempty" form:"commit"`
}
