package service

import (
	"jingdong/dao"
	"jingdong/models"
	"jingdong/utils"
)

func Register(c models.Register) (bool, string) {

	ok, err := utils.GetCk(c.Number, c.Code)
	if !ok {
		return false, err
	}
	ok, state := dao.Register(c.Username, c.Password, c.Number)
	return ok, state
}
