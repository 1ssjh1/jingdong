package utils

import (
	"JD/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetNews() models.News {

	clients := &http.Client{}
	urls := "https://we.cqupt.edu.cn/api/news/jw_list.php"
	reqs, _ := http.NewRequest("GET", urls, nil)
	var News models.News
	resps, _ := clients.Do(reqs)
	info, _ := ioutil.ReadAll(resps.Body)
	fmt.Println(resps.StatusCode)
	err := json.Unmarshal(info, &News)
	fmt.Println(err)
	//随机新闻 前端要求的
	var Rand models.News
	return Rand

}
