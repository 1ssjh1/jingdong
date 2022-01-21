package service

import (
	"Goto/models"
	"Goto/dao"
)

func Login(u models.Login) (bool,string) {
	ok,sate :=dao.Login(u)
	if ok {

	}
	return ok,sate

}
