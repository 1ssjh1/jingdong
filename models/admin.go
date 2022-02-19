package models

import "mime/multipart"

type AllInfo struct {
	All []All
}
type All struct {
	State string
	Order []OneOrder
}
type OneOrder struct {
	Oid   int
	Uid   int
	Gid   int
	Count int
}
type UpdateUserOrder struct {
	Oid   int    `json:"oid,omitempty" form:"oid"`
	State string `json:"state,omitempty" form:"state"`
}
type GoodsAdd struct {
	Image     *multipart.FileHeader `json:"image,omitempty" form:"image"`
	Gname     string                `json:"gname,omitempty" form:"gname"`
	Category  string                `json:"type,omitempty" form:"type"`
	Introduce string                `json:"introduce,omitempty" form:"introduce"`
	Price     float64               `json:"price,omitempty" form:"price"`
}
type UpdateGoods struct {
	GoodsAdd
	Gid int `json:"gid,omitempty" form:"gid"`
}
