package service

import (
	"jingdong/dao"
	"jingdong/models"
)

func BalanceGet(u string) (bool, string) {
	ok, state := dao.GetBalance(u)
	return ok, state

}
func BalanceCharge(u models.Balance) (bool, string) {
	ok, state := dao.ChargeBalance(u)
	return ok, state
}
