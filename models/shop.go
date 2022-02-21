package models

type AllShop struct {
	All []Type
}
type Type struct {
	Type  string
	Goods []Info
}
type UserAllGoods struct {
	All []Info
}
type Info struct {
	GoodsBasicInfo
	Goods
}

// Goods 商品信息
type Goods struct {
	Price     float64
	Sales     int
	Commit    int
	Grate     int
	Introduce string
}
type GoodsBasicInfo struct {
	Gid  int
	Name string
	Url  string
	Type string
}

// UpdateOrder 删除订单
type UpdateOrder struct {
	BasicInfo
	Oid int `json:"oid,omitempty" form:"oid"`
}
type AddChart struct {
	BasicInfo
	Gid   int `json:"gid,omitempty" form:"gid" `
	Count int `json:"count,omitempty" form:"count"`
}

// ShopChart 用于加入购车
type ShopChart struct {
	BasicInfo
	ChartId int `form:"chart_id" json:"chart_id,omitempty"`
	Count   int `form:"count" json:"count,omitempty"`
}
type Userinfo struct {
	BasicInfo
}

// AllChart  用于展示所有购物信息
type AllChart struct {
	BasicInfo
	ChartList []ChartShop
}

//后端生成订单所需信息
type ChartShop struct {
	ChartId int
	Gid     int
	Good    string
	Count   string
}

// 用户生成订单信息
type Order struct {
	BasicInfo
	ChartId []int `json:"chart_id,omitempty" form:"chart_id"`
}

//获取商品评论用户信息
type Commits struct {
	BasicInfo
	Gid string `json:"gid,omitempty" form:"gid"`
}

//所有评论
type AllCommit struct {
	Gid       int         `json:"gid,omitempty" form:"gid"`
	Introduce string      `json:"introduce,omitempty" form:"introduce"`
	Onecomit  []OneCommit `json:"onecomit,omitempty" form:"onecomit"`
}

//单条评论 用于 切片
type OneCommit struct {
	Oid    int    `json:"oid,omitempty" form:"oid"`
	Url    string `json:"url,omitempty" form:"url"`
	Commit string `json:"commit,omitempty" form:"commit"`
}
