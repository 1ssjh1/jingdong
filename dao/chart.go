package dao

import (
	"fmt"
	"jingdong/models"
)

func AddChart(chart models.ShopChart) (bool, string) {
	stm, err := DB.Prepare("select gid,cid from shop_chart where uid=?")
	if err != nil {
		fmt.Println(err)
		return false, ""
	}
	var Temple models.ShopChart
	rows, err := stm.Query(chart.Uid)
	for rows.Next() {
		rows.Scan(&Temple.Gid, &Temple.Cid)
		if Temple.Gid == chart.Gid && Temple.Cid == chart.Cid {
			return false, "已经加入购物车了 试试别的吧"
		}
	}

	stm, err = DB.Prepare(" insert into shop_chart(uid,gid,Cid,count) values(?,?,?,?)")
	if err != nil {
		fmt.Println(err)

		return false, ""
	}
	_, err = stm.Exec(chart.Uid, chart.Gid, chart.Cid, chart.Count)
	if err != nil {
		fmt.Println(err)

		return false, ""
	}
	return true, "你的宝贝已经躺在购物车里了哦"

}
func AllChart(user models.Userinfo) (bool, []models.AllChart) {
	stm, err := DB.Prepare("select uid from user_info where name=?")
	if err != nil {
		fmt.Println(err)
		return false, nil
	}
	var Temple models.Userinfo
	err = stm.QueryRow(user.Username).Scan(&Temple.Uid)
	if err != nil {
		fmt.Println(err)
		return false, nil
	}
	if Temple.Uid != user.Uid {
		fmt.Println("参数错误")
		return false, nil
	}

	all := make([]models.AllChart, 1, 1)
	all[0].Username = user.Username
	all[0].Uid = Temple.Uid
	stm, err = DB.Prepare("select chart_id,gid, cid,count from shop_chart where uid=? ")
	if err != nil {
		fmt.Println(err)
		return false, nil
	}
	rows, err := stm.Query(user.Uid)
	if err != nil {
		fmt.Println(err)
		return false, nil
	}
	Templevalue := models.ChartShop{}
	for rows.Next() {
		rows.Scan(&Templevalue.ChartId, &Templevalue.Gid, &Templevalue.Cid, &Templevalue.Count)
		stm, err = DB.Prepare("select name from goods_list where Gid=? ")
		err = stm.QueryRow(Templevalue.Gid).Scan(&Templevalue.Good)
		stm, err = DB.Prepare("select type from goods_type where Cid=?")
		err = stm.QueryRow(Templevalue.Cid).Scan(&Templevalue.Types)
		all[0].ChartList = append(all[0].ChartList, Templevalue)
	}
	return true, all

}
func UpdateChart(chart models.ShopChart) (bool, string) {

	//将Count设置为0 就意味着删除
	if chart.Count == 0 {

		//查看数据是否合法
		stm, err := DB.Prepare("select uid  from shop_chart where gid=? and cid=?")
		var Temple models.ShopChart
		err = stm.QueryRow(chart.Gid, chart.Cid).Scan(&Temple.Uid)
		if err != nil {
			fmt.Println(err)
			return false, "失败"
		}
		if Temple.Uid != chart.Uid {
			return false, "用户信息不匹配"
		}

		stm, err = DB.Prepare("delete from shop_chart where gid=? and cid=?")
		if err != nil {
			return false, "失败"

		}
		_, err = stm.Exec(chart.Gid, chart.Cid)
		if err != nil {
			return false, "失败"

		}
		return true, "宝贝忍痛离开了购物车"

	}
	stm, err := DB.Prepare("update shop_chart set cid=?,count =? where uid=? and gid=? and cid=? ")
	if err != nil {
		fmt.Println(err)
		return false, "失败"
	}
	_, err = stm.Exec(chart.Cid, chart.Count, chart.Uid, chart.Gid, chart.Cid)
	if err != nil {
		return false, "失败"
	}
	return true, "操作成功"
}
