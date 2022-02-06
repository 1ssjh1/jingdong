package models

type UpdateOrder struct {
	Uid   int    `json:"uid,omitempty" form:"uid"`
	Token string `json:"token,omitempty" form:"token"`
	Oid   string `json:"oid,omitempty" form:"oid"`
}
