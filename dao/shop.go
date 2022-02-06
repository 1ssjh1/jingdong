package dao

import (
	"fmt"
	"jingdong/models"
)

func AllShops() []models.Goods {
	stm, err := DB.Prepare("select  Gid ,name  from goods_list ")
	if err != nil {
		fmt.Println(err, "  1")
		return nil
	}
	defer stm.Close()
	AllShop := make([]models.Goods, 0, 1)

	var TempleShop models.Goods
	count := 0
	rows, err := stm.Query()
	if err != nil {
		fmt.Println(err, "err")
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&TempleShop.Gid, &TempleShop.Name)
		AllShop = append(AllShop, TempleShop)
		if err != nil {
			fmt.Println(err)
		}
		count++

	}
	fmt.Println(AllShop)
	if rows != nil {
	}
	stm, err = DB.Prepare("select sales,commit,Grate,introduce from goods_info where Gid=?")
	if err != nil {
		fmt.Println(err, 1)
	}
	for i := 0; i < count; i++ {
		rows, err = stm.Query(AllShop[i].Gid)
		if err != nil {
			break

			return nil

		}
		for rows.Next() {
			err = rows.Scan(&AllShop[i].Sales, &AllShop[i].Commit, &AllShop[i].Grate, &AllShop[i].Introduce)
			if err != nil {
				fmt.Println(err)
				return nil
			}
		}
	}
	var TempleType models.Choose
	stm, err = DB.Prepare("select Cid,type ,price from goods_type where Gid=?")
	for i := 0; i < count; i++ {
		rows, err = stm.Query(AllShop[i].Gid)
		var n = 0
		for rows.Next() {
			err = rows.Scan(&TempleType.Cid, &TempleType.Types, &TempleType.Price)
			AllShop[i].Choose = append(AllShop[i].Choose, TempleType)
			if err != nil {
				return nil
			}
			n++
		}
	}
	return AllShop
}
func MakeOrder(order models.Order) (bool, string) {

	//检验数据是否合法
	stm, err := DB.Prepare("select uid from shop_chart where chart_id =?")
	if err != nil {
		return false, "错误"
	}
	var Temple models.Order
	for _, v := range order.ChartId {
		rows, err := stm.Query(v)
		if err != nil {
			return false, "错误"
		}
		for rows.Next() {
			rows.Scan(&Temple.Uid)
			if Temple.Uid != order.Uid {
				return false, "错误"

			}
		}
	}
	tx, err := DB.Begin()
	if err != nil {
		tx.Rollback()
		return false, "错误"
	}
	stm, err = tx.Prepare("insert into user_order(oid,uid,gid,cid,count) values (?, ?,?,?,?)")
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return false, "订单创建失败"

	}
	var Temp models.ChartShop
	allinfo := make([]models.ChartShop, 0, 1)
	stm1, err := tx.Prepare("select uid ,gid ,cid ,Count from shop_chart where chart_id =?")
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return false, "订单创建失败"

	}
	for _, value := range order.ChartId {
		err = stm1.QueryRow(value).Scan(&Temp.Uid, &Temp.Gid, &Temp.Cid, &Temp.Count)
		if err != nil {
			fmt.Println(err)
			tx.Rollback()
			return false, "订单创建失败"

		}
		_, err = stm.Exec(value, Temp.Uid, Temp.Gid, Temp.Cid, Temp.Count)
		allinfo = append(allinfo, Temp)
		if err != nil {
			fmt.Println(err)
			tx.Rollback()
			return false, "订单创建失败"

		}
	}
	stm, err = tx.Prepare("delete from shop_chart where chart_id=?")
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return false, "订单创建失败"

	}
	for _, v := range order.ChartId {
		_, err = stm.Exec(v)
		if err != nil {

			fmt.Println(err)
			tx.Rollback()
			return false, "订单创建失败"
		}
	}
	//后续补上扣款功能 不扣款咋能行
	stm, err = tx.Prepare("select price from goods_type where Cid=? and Gid=?")
	if err != nil {

		fmt.Println(err, "4")
		tx.Rollback()
		return false, "订单创建失败"
	}
	var allprice = 0
	var tempv int
	for _, v := range allinfo {
		fmt.Println(v)
		err = stm.QueryRow(v.Cid, v.Gid).Scan(&tempv)
		if err != nil {
			fmt.Println(err, "3")
			tx.Rollback()
			return false, "订单创建失败"
		}
		allprice += tempv
	}
	//扣款
	stm, err = tx.Prepare("select balance from user_info where uid=?")
	if err != nil {

		fmt.Println(err, "2")
		tx.Rollback()
		return false, "订单创建失败"
	}
	err = stm.QueryRow(order.Uid).Scan(&tempv)
	if err != nil {

		fmt.Println(err, "1")
		tx.Rollback()
		return false, "订单创建失败"
	}
	if tempv < allprice {
		tx.Rollback()
		return false, "你个穷逼"
	}
	stm, err = tx.Prepare("update user_info set balance =? where uid=?")
	if err != nil {

		fmt.Println(err)
		tx.Rollback()
		return false, "订单创建失败"
	}
	_, err = stm.Exec(tempv-allprice, order.Uid)
	if err != nil {

		fmt.Println(err, "5")
		tx.Rollback()
		return false, "订单创建失败"
	}
	tx.Commit()
	return true, "订单提交成功"
}
func AllOrder(user models.User) (bool, *models.UserOrder) {
	stm, err := DB.Prepare("select uid from user_info where name =?")
	if err != nil {
		return false, nil
	}
	var Uid int
	err = stm.QueryRow(user.Username).Scan(&Uid)
	if err != nil {
		return false, nil
	}
	stm, err = DB.Prepare("select oid,state,gid ,cid ,count from user_order where uid=?")
	if err != nil {
		return false, nil
	}
	row, err := stm.Query(Uid)
	if err != nil {
		return false, nil
	}
	var temple models.AllOrder

	var all models.UserOrder
	all.Uid = Uid
	for row.Next() {
		err := row.Scan(&temple.Oid, &temple.State, &temple.Gid, &temple.Cid, &temple.Count)
		fmt.Println(err)
		all.Allorder = append(all.Allorder, temple)
	}
	return true, &all
}
func UpdateOrder(order models.UpdateOrder) (bool, string) {
	stm, err := DB.Prepare("select state from user_order where oid =?")
	if err != nil {
		fmt.Println(err)
		return false, "失败"
	}
	var Orderstate string
	err = stm.QueryRow(order.Uid).Scan(&Orderstate)
	if Orderstate == "未发货" {
		return false, "还没发货呢 着啥急"
	}
	stm, err = DB.Prepare("update user_order set state=? where oid =?")
	if err != nil {
		return false, "错误"
	}
	_, err = stm.Exec("已完成", order.Oid)
	if err != nil {
		fmt.Println(err)

		return false, "失败"
	}
	return true, "确认收获成功"

}
func DeleteOrder(order models.UpdateOrder) (bool, string) {
	stm, err := DB.Prepare("select uid,state from user_order where oid=?")
	if err != nil {
		return false, "失败"
	}
	var (
		templeuid   int
		templestate string
	)
	err = stm.QueryRow(order.Oid).Scan(&templeuid, &templestate)
	if err != nil {
		return false, "失败"
	}
	if templeuid != order.Uid {
		return false, "怎么可以动别人的订单呢"
	}
	if templestate == "已完成" {
		return false, "已经完成订单了 概不负责了哦"
	}
	stm, err = DB.Prepare("delete from user_order where oid =?")
	if err != nil {
		return false, "失败"
	}
	_, err = stm.Exec(order.Oid)
	if err != nil {
		return false, "失败"
	}
	return true, "订单销毁成功"
}
func Commit(commit models.Commit) (bool, string) {
	stm, err := DB.Prepare("select uid  ,state,gid,cid from user_order where oid =?")
	if err != nil {
		fmt.Println(err)
		return false, "评论提交失败"
	}
	var (
		uid   int
		state string
		gid   int
		cid   int
	)
	err = stm.QueryRow(commit.Oid).Scan(&uid, &state, &gid, &cid)
	if err != nil {
		fmt.Println(err)
		return false, "评论提交失败"
	}
	if uid != commit.Uid || state != "已完成" {
		return false, "商品状态错误"
	}
	stm, err = DB.Prepare("insert into goods_commit (Gid,oid,cid,commit) values (?,?,?,?)")
	if err != nil {
		fmt.Println(err)
		return false, "评论提交失败"
	}
	_, err = stm.Exec(gid, commit.Oid, cid, commit.Commit)
	if err != nil {
		fmt.Println(err)
		return false, "评论提交失败"
	}
	return true, "评论提价成功"
}
func GetCommit(commit models.Commits) *models.AllCommit {
	stm, err := DB.Prepare("select introduce from goods_info where GId =?")
	if err != nil {
		return nil
	}
	var allcommit models.AllCommit
	err = stm.QueryRow(commit.Gid).Scan(&allcommit.Introduce)
	if err != nil {
		return nil
	}
	stm, err = DB.Prepare("select oid ,cid ,commit from goods_commit where Gid=?")
	if err != nil {
		return nil
	}
	var onecommit models.OneCommit
	row, err := stm.Query(commit.Gid)
	if err != nil {
		return nil
	}
	for row.Next() {
		row.Scan(&onecommit.Oid, &onecommit.Cid, &onecommit.Commit)
		allcommit.Onecomit = append(allcommit.Onecomit, onecommit)
	}
	return &allcommit
}
