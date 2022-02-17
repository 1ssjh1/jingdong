package dao

import (
	"JD/models"
	"errors"
	"fmt"
	"sync"
)

func AllShops() *models.UserAllGoods {
	var (
		AllShop         models.UserAllGoods
		templebasicinfo models.Info
	)
	stm, err := DB.Prepare("select Gid ,name ,url,type from goods_list")
	if err != nil {
		return nil
	}
	rows, err := stm.Query()
	if err != nil {
		return nil
	}
	go func() {
		for rows.Next() {
			err := rows.Scan(&templebasicinfo.Gid, &templebasicinfo.Name, &templebasicinfo.Url, &templebasicinfo.Type)
			if err != nil {
				fmt.Println(err, "fw")

			}
			templebasicinfo.Url = "https://sanser.ltd/static/" + templebasicinfo.Url
			AllShop.All = append(AllShop.All, templebasicinfo)
		}
	}()
	//for rows.Next() {
	//	err := rows.Scan(&templebasicinfo.Gid, &templebasicinfo.Name, &templebasicinfo.Url, &templebasicinfo.Type)
	//	if err != nil {
	//		fmt.Println(err, "fw")
	//		return nil
	//	}
	//	templebasicinfo.Url = "https://sanser.ltd/static/" + templebasicinfo.Url
	//	AllShop.All = append(AllShop.All, templebasicinfo)
	//}
	fmt.Println(AllShop)
	s := sync.WaitGroup{}
	stms, err := DB.Prepare("select sales,commit,Grate ,introduce ,price from goods_info where Gid =?")
	chanel_1 := make(chan int, len(AllShop.All))
	chanel_2 := make(chan int, len(AllShop.All))

	for k, v := range AllShop.All {
		chanel_1 <- k
		chanel_2 <- v.Gid

	}

	//蹩脚 并发
	s.Add(len(AllShop.All))
	for i := 0; i <= cap(AllShop.All); i++ {
		go func() {
			v := <-chanel_2
			k := <-chanel_1
			row, err := stms.Query(v)
			for row.Next() {
				temple := models.Goods{}
				err = row.Scan(&temple.Sales, &temple.Commit, &temple.Grate, &temple.Introduce, &temple.Price)
				AllShop.All[k].Goods = temple
				if err != nil {
					fmt.Println(err, "ger")
				}
				s.Done()
			}
		}()

	}

	s.Wait()
	close(chanel_1)
	close(chanel_2)
	return &AllShop

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
	stm, err = tx.Prepare("insert into user_order(uid,gid,count) values (?,?,?)")
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return false, "订单创建失败"
	}
	var Temp models.ChartShop
	allinfo := make([]models.ChartShop, 0, 1)
	stm1, err := tx.Prepare("select gid ,Count from shop_chart where chart_id =?")
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return false, "订单创建失败"

	}
	for _, value := range order.ChartId {
		err = stm1.QueryRow(value).Scan(&Temp.Gid, &Temp.Count)
		if err != nil {
			fmt.Println(err)
			tx.Rollback()
			return false, "订单创建失败"

		}
		_, err = stm.Exec(order.Uid, Temp.Gid, Temp.Count)
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
	stm, err = tx.Prepare("select price from goods_type where  Gid=?")
	if err != nil {

		fmt.Println(err, "4")
		tx.Rollback()
		return false, "订单创建失败"
	}
	var allprice = 0
	var tempv int
	for _, v := range allinfo {
		fmt.Println(v)
		err = stm.QueryRow(v.Gid).Scan(&tempv)
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
	//stm, err := DB.Prepare("select uid from user_info where name =?")
	//if err != nil {
	//	return false, nil
	//}
	//var Uid int
	//err = stm.QueryRow(user.Username).Scan(&Uid)
	//if err != nil {
	//	return false, nil
	//}
	stm, err := DB.Prepare("select oid,state,gid ,count from user_order where uid=?")
	if err != nil {
		return false, nil
	}
	row, err := stm.Query(user.Uid)
	if err != nil {
		return false, nil
	}
	var temple models.AllOrder

	var all models.UserOrder
	all.BasicInfo = user.BasicInfo
	for row.Next() {
		err := row.Scan(&temple.Oid, &temple.State, &temple.Gid, &temple.Count)
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
	stm, err = DB.Prepare("select oid ,url,commit from goods_commit where Gid=?")
	if err != nil {
		return nil
	}
	var onecommit models.OneCommit
	row, err := stm.Query(commit.Gid)
	if err != nil {
		return nil
	}
	for row.Next() {
		row.Scan(&onecommit.Oid, &onecommit.Url, &onecommit.Commit)
		onecommit.Url = "https://sanser.ltd/static/" + onecommit.Url
		allcommit.Onecomit = append(allcommit.Onecomit, onecommit)
	}
	return &allcommit
}
func Class() (*models.AllShop, error) {
	stm, err := DB.Prepare("select type from goods_class")
	if err != nil {
		fmt.Println(err, "1")

		err = errors.New("查询失败")
		return nil, err
	}
	var all models.AllShop
	row, err := stm.Query()
	if err != nil {
		fmt.Println(err)

		err = errors.New("查询失败")
		return nil, err
	}
	var onetype models.Type
	for row.Next() {
		err = row.Scan(&onetype.Type)
		if err != nil {
			fmt.Println(err)

			err = errors.New("查询失败")
			return nil, err
		}
		all.All = append(all.All, onetype)
	}
	var (
		chanel_1 = make(chan string, len(all.All))
		chanel_2 = make(chan int, len(all.All))
	)
	s := sync.WaitGroup{}
	stm, err = DB.Prepare("select Gid ,name, url from goods_list where type =?")
	for k, v := range all.All {
		s.Add(1)
		chanel_1 <- v.Type
		chanel_2 <- k
	}
	for i := 0; i <= len(all.All); i++ {
		go func() {
			Onetype := <-chanel_1
			n := <-chanel_2
			row, err := stm.Query(Onetype)
			if err != nil {
				err = errors.New("查询失败")
			}
			var temple models.Info
			for row.Next() {
				err = row.Scan(&temple.Gid, &temple.Name, &temple.Url)
				if err != nil {
					fmt.Println(err)
					err = errors.New("查询失败")
				}
				fmt.Println(temple)
				temple.Url = "https://sanser.ltd/static/" + temple.Url

				all.All[n].Goods = append(all.All[n].Goods, temple)

				fmt.Println(all.All[n].Goods[0])
			}
			s.Done()

		}()
	}
	s.Wait()

	stm, err = DB.Prepare("select sales,commit ,Grate,introduce,price from goods_info where Gid =?")
	for k, v := range all.All {
		//双层查找属实很蹩脚 n gid
		for m, n := range v.Goods {
			var temple models.Goods
			row, err = stm.Query(n.Gid)
			if err != nil {
				fmt.Println(err)

				err = errors.New("查询失败")
				return nil, err
			}
			for row.Next() {
				err = row.Scan(&temple.Sales, &temple.Commit, &temple.Grate, &temple.Introduce, &temple.Price)
				if err != nil {
					fmt.Println(err)
					err = errors.New("查询失败")
					return nil, err
				}
				fmt.Println(temple)
				all.All[k].Goods[m].Goods = temple
			}
		}
	}
	fmt.Println("m", all.All)
	return &all, nil

}
