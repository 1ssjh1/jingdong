package dao

import (
	"JD/models"
	"errors"
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

			}
			templebasicinfo.Url = "https://sanser.ltd/static/" + templebasicinfo.Url
			AllShop.All = append(AllShop.All, templebasicinfo)
		}
	}()
	// 嘿嘿 开始蹩脚并发
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
		tx.Rollback()
		return false, "订单创建失败"
	}
	var Temp models.ChartShop
	allinfo := make([]models.ChartShop, 0, 1)
	stm1, err := tx.Prepare("select gid ,Count from shop_chart where chart_id =?")
	if err != nil {
		tx.Rollback()
		return false, "订单创建失败"

	}
	for _, value := range order.ChartId {
		err = stm1.QueryRow(value).Scan(&Temp.Gid, &Temp.Count)
		if err != nil {
			tx.Rollback()
			return false, "订单创建失败"

		}
		_, err = stm.Exec(order.Uid, Temp.Gid, Temp.Count)
		allinfo = append(allinfo, Temp)
		if err != nil {
			tx.Rollback()
			return false, "订单创建失败"

		}
	}
	stm, err = tx.Prepare("delete from shop_chart where chart_id=?")
	if err != nil {
		tx.Rollback()
		return false, "订单创建失败"

	}
	for _, v := range order.ChartId {
		_, err = stm.Exec(v)
		if err != nil {

			tx.Rollback()
			return false, "订单创建失败"
		}
	}
	//后续补上扣款功能 不扣款咋能行
	stm, err = tx.Prepare("select price from goods_info where  Gid=?")
	if err != nil {

		tx.Rollback()
		return false, "订单创建失败"
	}
	var allprice float64
	var tempv float64
	for _, v := range allinfo {
		err = stm.QueryRow(v.Gid).Scan(&tempv)
		if err != nil {
			tx.Rollback()
			return false, "订单创建失败"
		}
		allprice += tempv
	}
	//扣款
	stm, err = tx.Prepare("select balance from user_info where uid=?")
	if err != nil {

		tx.Rollback()
		return false, "订单创建失败"
	}
	err = stm.QueryRow(order.Uid).Scan(&tempv)
	if err != nil {

		tx.Rollback()
		return false, "订单创建失败"
	}
	if tempv < allprice {
		tx.Rollback()
		return false, "你个穷逼"
	}
	stm, err = tx.Prepare("update user_info set balance =? where uid=?")
	if err != nil {

		tx.Rollback()
		return false, "订单创建失败"
	}
	_, err = stm.Exec(tempv-allprice, order.Uid)
	if err != nil {

		tx.Rollback()
		return false, "订单创建失败"
	}
	stm, err = tx.Prepare("update goods_info set sales=sales+1 where GId=?")
	if err != nil {

		tx.Rollback()
		return false, "订单创建失败"
	}
	_, err = stm.Exec(Temp.Gid)
	if err != nil {

		tx.Rollback()
		return false, "订单创建失败"
	}
	tx.Commit()
	return true, "订单提交成功"
}
func AllOrder(user models.User) (bool, *models.UserOrder) {

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
		row.Scan(&temple.Oid, &temple.State, &temple.Gid, &temple.Count)
		all.Allorder = append(all.Allorder, temple)
	}
	return true, &all
}
func UpdateOrder(order models.UpdateOrder) error {
	stm, err := DB.Prepare("select state from user_order where oid =?")
	if err != nil {
		err = errors.New("失败")
		return err
	}
	var Orderstate string
	err = stm.QueryRow(order.Oid).Scan(&Orderstate)
	if Orderstate != "已支付" {
		if Orderstate == "已完成" {
			err = errors.New("订单已经完成了哦")
			return err
		}
		err = errors.New("还未发货 着啥急")
		return err
	}
	stm, err = DB.Prepare("update user_order set state=? where oid =?")
	if err != nil {
		err = errors.New("失败")
		return err
	}
	_, err = stm.Exec("已完成", order.Oid)
	if err != nil {
		err = errors.New("失败")
		return err
	}
	return nil

}
func DeleteOrder(order models.UpdateOrder) error {
	stm, err := DB.Prepare("select uid,state from user_order where oid=?")
	if err != nil {
		err = errors.New("失败")
		//return false, "失败"
		return err

	}
	var (
		templeuid   int
		templestate string
	)
	err = stm.QueryRow(order.Oid).Scan(&templeuid, &templestate)
	if err != nil {
		err = errors.New("失败")
		//return false, "失败"
		return err
	}
	if templeuid != order.Uid {
		err = errors.New("怎么可以动别人的订单呢")
		return err
	}
	if templestate != "已完成" {

		err = errors.New("客官 订单还没跑完呢 不要抛弃我呀")
		return err
	}
	stm, err = DB.Prepare("delete from user_order where oid =?")
	if err != nil {
		err = errors.New("失败")
		//return false, "失败"
		return err
		//return false, "失败"
	}
	_, err = stm.Exec(order.Oid)
	if err != nil {
		err = errors.New("失败")
		//return false, "失败"
		return err
	}
	return nil
}
func Commit(commit models.Commit) (bool, string) {
	stm, err := DB.Prepare("select uid ,state,gid from user_order where oid =?")
	if err != nil {
		return false, "评论提交失败"
	}
	var (
		uid   int
		state string
		gid   int
	)
	err = stm.QueryRow(commit.Oid).Scan(&uid, &state, &gid)
	if err != nil {
		return false, "评论提交失败"
	}
	if uid != commit.Uid || state != "已完成" {
		return false, "商品状态错误"
	}
	stm, err = DB.Prepare("select cid from goods_commit where oid=?")
	if err != nil {
		return false, "评论提交失败"
	}
	var commitID int
	stm.QueryRow(commit.Oid).Scan(&commitID)

	if commitID != 0 {
		return false, "您已经评论过了"
	}
	stm, err = DB.Prepare("insert into goods_commit (Gid,oid,commit) values (?,?,?)")
	if err != nil {
		return false, "评论提交失败"
	}
	_, err = stm.Exec(gid, commit.Oid, commit.Commit)
	if err != nil {
		return false, "评论提交失败"
	}
	stm, err = DB.Prepare("update goods_info set commit=commit+1 where GId=?")
	if err != nil {
		return false, "评论提交失败"
	}
	_, err = stm.Exec(gid)
	if err != nil {
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

		//发现用户评论的时候并没有附带上 图片 这里面的url 就弄个默认值吧 （doge
		row.Scan(&onecommit.Oid, &onecommit.Url, &onecommit.Commit)
		onecommit.Url = "https://sanser.ltd/static/" + onecommit.Url
		allcommit.Onecomit = append(allcommit.Onecomit, onecommit)
	}
	return &allcommit
}
func Class() (*models.AllShop, error) {
	stm, err := DB.Prepare("select type from goods_class")
	if err != nil {

		err = errors.New("查询失败")
		return nil, err
	}
	var all models.AllShop
	row, err := stm.Query()
	if err != nil {

		err = errors.New("查询失败")
		return nil, err
	}
	var onetype models.Type
	for row.Next() {
		err = row.Scan(&onetype.Type)
		if err != nil {

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
					err = errors.New("查询失败")
				}
				temple.Url = "https://sanser.ltd/static/" + temple.Url

				all.All[n].Goods = append(all.All[n].Goods, temple)

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

				err = errors.New("查询失败")
				return nil, err
			}
			for row.Next() {
				err = row.Scan(&temple.Sales, &temple.Commit, &temple.Grate, &temple.Introduce, &temple.Price)
				if err != nil {
					err = errors.New("查询失败")
					return nil, err
				}
				all.All[k].Goods[m].Goods = temple
			}
		}
	}
	return &all, nil

}
