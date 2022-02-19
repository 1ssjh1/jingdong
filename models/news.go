package models

type News struct {
	Status  int    `json:"status"`
	Message []Data `json:"data"`
	//Data
}
type Data struct {
	Articleid int    `json:"articleid"`
	Title     string `json:"title"`
}
