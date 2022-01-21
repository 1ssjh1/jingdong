package models

import "time"

type Register struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Code     string  `form:"code" json:"code" binding:"required"`
	Number string `form:"number" json:"number" binding:"required"`
}
type Login struct {
	Username string `form:"username" json:"username,omitempty" binding:"required"`
	Password string `form:"password" json:"password,omitempty" binding:"required"`
}
type Message struct {
	Send string
	Recive string
	Sendtime time.Time

}