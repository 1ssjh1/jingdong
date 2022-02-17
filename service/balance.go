package service

import (
	"JD/dao"
	"JD/models"
)

func BalanceGet(u string) (bool, interface{}) {
	ok, state := dao.GetBalance(u)
	return ok, state

}
func BalanceCharge(u models.Balance) (bool, string) {
	ok, state := dao.ChargeBalance(u)
	return ok, state
}
