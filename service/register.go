package service

import (
	"JD/dao"
	"JD/models"
	"JD/utils"
)

func Register(c models.Register) (bool, error) {

	ok, err := utils.GetCk(c.Number, c.Code)
	if !ok {
		return false, err
	}
	ok, state := dao.Register(c.Username, c.Password, c.Number)
	return ok, state
}
