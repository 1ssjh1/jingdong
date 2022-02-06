package dao

import (
	"jingdong/models"
)

func GetBalance(u string) (bool, string) {
	stm, err := DB.Prepare("select  balance from user_info where name = ?")
	if err != nil {
		return false, "查找失败"
	}
	defer stm.Close()
	rows, err := stm.Query(u)
	if err != nil {
		return false, "查找失败"
	}
	var tmp string
	for rows.Next() {
		rows.Scan(&tmp)
	}
	if tmp == "" {
		return false, "未查到用户信息"
	}
	return true, tmp

}
func ChargeBalance(u models.Balance) (bool, string) {

	stm, err := DB.Prepare("update user_info set balance=? where name=?")
	if err != nil {
		return false, "充值失败"
	}
	_, err = stm.Exec(u.Balance, u.Username)
	if err != nil {
		return false, "充值失败"
	}
	return true, "充值成功"
}
