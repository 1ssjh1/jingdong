package utils

import (
	"JD/models"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

func GetNews() (*models.News, error) {

	clients := &http.Client{}
	urls := "https://we.cqupt.edu.cn/api/news/jw_list.php"
	req, _ := http.NewRequest("GET", urls, nil)
	var News models.News
	resp, _ := clients.Do(req)
	info, _ := ioutil.ReadAll(resp.Body)
	err := json.Unmarshal(info, &News)
	if err != nil {
		return nil, err
	}
	//随机新闻 前端要求的 emmm
	defer resp.Body.Close()
	var Rand models.News
	var wait sync.WaitGroup
	wait.Add(1)

	rand.Seed(time.Now().UnixNano())
	go func() {
		var count = 0
		for true {
		loop:
			t := rand.Intn(20)
			for _, v := range Rand.Message {
				if v.Title == News.Message[t].Title {
					goto loop
				}
			}
			Rand.Message = append(Rand.Message, News.Message[t])
			count++
			if count == 4 {
				wait.Done()
				break
			}
		}
	}()
	wait.Wait()
	Rand.Status = News.Status
	return &Rand, nil

}
