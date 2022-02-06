package models

type Goods struct {
	Gid       int
	Name      string
	Sales     int
	Commit    int
	Grate     int
	Introduce string
	Choose    []Choose
}
type Choose struct {
	Cid   int
	Types string
	Price int
}

type Chart struct {
}

// ShopChart 用于加入购车 以及对购物车进行修改
type ShopChart struct {
	Username string `form:"username" json:"username,omitempty"`
	Uid      int    `form:"uid" json:"uid,omitempty"`
	Gid      int    `form:"gid" json:"gid,omitempty"`
	Cid      int    `form:"cid" json:"cid,omitempty"`
	Count    int    `form:"count" json:"count,omitempty"`
	Token    string `json:"token,omitempty" form:"token"`
}
type Userinfo struct {
	Uid      int    `form:"uid" json:"uid,omitempty"`
	Username string `form:"username" json:"username,omitempty"`
	Token    string `json:"token ,omitempty" form:"token"`
}

// AllChart  用于展示所有购物信息
type AllChart struct {
	Username  string
	Uid       int
	ChartList []ChartShop
}
type ChartShop struct {
	Uid     int
	ChartId int
	Gid     int
	Cid     int
	Count   string
	Good    string
	Types   string
}
type Order struct {
	Username string `json:"username ,omitempty" form:"username"`
	Uid      int    `json:"uid,omitempty" form:"uid"`
	ChartId  []int  `json:"chart_id,omitempty" form:"chart_id"`
	Token    string `json:"token,omitempty" form:"token"`
}
type Commits struct {
	Uid      int    `json:"uid,omitempty" form:"uid"`
	Username string `json:"username,omitempty" form:"username"`
	Gid      int    `json:"gid,omitempty" form:"gid"`
	Token    string `json:"token,omitempty" form:"token"`
}
type AllCommit struct {
	Gid       int         `json:"gid,omitempty" form:"gid"`
	Introduce string      `json:"introduce,omitempty" form:"introduce"`
	Onecomit  []OneCommit `json:"onecomit,omitempty" form:"onecomit"`
}
type OneCommit struct {
	Oid    int    `json:"oid,omitempty" form:"oid"`
	Cid    int    `json:"cid,omitempty" form:"cid"`
	Commit string `json:"commit,omitempty" form:"commit"`
}
