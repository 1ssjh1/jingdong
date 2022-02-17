package models

import (
	"mime/multipart"
)

// BasicInfo 用户基准信息 一般用于 token解析数据绑定
type BasicInfo struct {
	Uid      int
	Username string
}

// Balance 用户余额充值
type Balance struct {
	BasicInfo
	Balance int `json:"balance,omitempty" form:"balance"`
}

// User token 附带用户信息
type User struct {
	BasicInfo
}

// UserOrder 用户订单返回
type UserOrder struct {
	BasicInfo
	Allorder []AllOrder
}

// AllOrder 单条订单

// Commit 用户提交评论
type Commit struct {
	BasicInfo
	Oid    int    `json:"oid,omitempty" form:"oid"`
	Commit string `json:"commit,omitempty" form:"commit"`
}

// UserImage 用户头像上传
type UserImage struct {
	BasicInfo
	Image *multipart.FileHeader `json:"image,omitempty" form:"image"`
}

// Info 用于用户个人信息 订单分类展示

type MyInfo struct {
	BasicInfo
	Balance  int
	ImageUrl string
	Category []Category
}
type AllOrder struct {
	Gid   int
	Oid   int
	Count int
	State string
}

type Category struct {
	State string
	Order []AllOrder
}
