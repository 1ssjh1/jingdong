package dao

import (
	"JD/models"
	"JD/utils"
	"os"
	// "JD/utils"
	// "JD/utils"
	// "JD/utils"
	"database/sql"
	"errors"

	"github.com/gin-gonic/gin"
)

func GetAllOrder() (*models.AllInfo, error) {
	var AllInfo models.AllInfo
	stm, err := DB.Prepare("select State from Order_state_type")
	if err != nil {
		err = errors.New("数据查询失败")
		return nil, err
	}
	row, err := stm.Query()
	if err != nil {

		err = errors.New("数据查询失败")
		return nil, err
	}
	var templeAll models.All
	for row.Next() {
		err := row.Scan(&templeAll.State)
		if err != nil {

			if err == sql.ErrNoRows {
				//do nothing
			} else {
				err = errors.New("数据查询失败")

				return nil, err
			}
		}
		AllInfo.All = append(AllInfo.All, templeAll)
	}
	var OneOrder models.OneOrder
	stm, err = DB.Prepare("select oid ,uid ,gid ,count from user_order where state=?")

	if err != nil {
		err = errors.New("数据查询失败")
		return nil, err
	}
	for key, value := range AllInfo.All {
		row, err = stm.Query(value.State)
		if err != nil {

			err = errors.New("数据查询失败")
			return nil, err
		}
		for row.Next() {
			err = row.Scan(&OneOrder.Oid, &OneOrder.Uid, &OneOrder.Gid, &OneOrder.Count)
			if err != nil {

				err = errors.New("数据查询失败")
				return nil, err
			}
			AllInfo.All[key].Order = append(AllInfo.All[key].Order, OneOrder)
		}
	}
	return &AllInfo, nil
}
func OrderChange(order models.UpdateUserOrder) (string, error) {
	stm, err := DB.Prepare("update user_order set state =? where oid=?")
	if err != nil {
		err = errors.New("操作失败")
		return "", err
	}
	_, err = stm.Exec(order.State, order.Oid)
	if err != nil {
		err = errors.New("操作失败")
		return "", err
	}
	return "操作成功", nil
}
func DeleteUserOrder(update models.UpdateUserOrder) (string, error) {
	stm, err := DB.Prepare("select state from user_order where oid=?")
	if err != nil {
		err = errors.New("用户信息错误")
		return "", err
	}
	var state string
	err = stm.QueryRow(update.Oid).Scan(&state)
	if err != nil {
		if err == sql.ErrNoRows {
			err = errors.New("订单不存在")
			return "", err
		}
		err = errors.New("操作失败")
		return "", err
	}

	stm, err = DB.Prepare("delete from user_order where oid =?")
	if err != nil {
		err = errors.New("用户信息错误")
		return "", err
	}
	_, err = stm.Exec(update.Oid)
	if err != nil {
		err = errors.New("操作失败")
		return "", err
	}
	return "操作成功", nil

}
func AddGoods(goods models.GoodsAdd, c *gin.Context) (string, error) {
	tx, err := DB.Begin()
	if err != nil {

		err = errors.New("商品添添加失败 请重试")
		return "", err
	}
	stm, err := tx.Prepare("insert into goods_info(introduce,price) values (?,?)")
	if err != nil {

		err = errors.New("商品添添加失败 请重试")
		tx.Rollback()
		return "", err
		//err=errors.New("")
	}
	_, err = stm.Exec(goods.Introduce, goods.Price)
	if err != nil {
		err = errors.New("商品添添加失败 请重试")
		return "", err
	}
	url, err := utils.SaveFile(goods.Image, c)
	if err != nil {

		//err = errors.New("商品添添加失败 请重试")
		tx.Rollback()
		return "", err
	}
	stm, err = tx.Prepare("insert into goods_list(name,url,type) values (?,?,?)")
	if err != nil {

		err = errors.New("商品添添加失败 请重试")
		return "", err
	}
	_, err = stm.Exec(goods.Gname, url, goods.Category)
	if err != nil {

		tx.Rollback()
		err = errors.New("商品添添加失败 请重试")
		return "", err
	}
	tx.Commit()
	return "商品添加成功", nil

}
func UpdateGoods(goods models.UpdateGoods, c *gin.Context) (string, error) {
	tx, err := DB.Begin()
	if err != nil {
		err = errors.New("数据更新失败 请稍后重试")
		return "", err
	}
	if goods.Image != nil {
		stm, err := DB.Prepare("select url from goods_list where Gid=?")
		if err != nil {

			err = errors.New("数据更新失败 请稍后重试")
			return "", err
		}
		var ToDelete string
		err = stm.QueryRow(goods.Gid).Scan(&ToDelete)
		if err != nil {

			err = errors.New("数据更新失败 请稍后重试")
			return "", err
		}
		err = os.Remove("/www/static/" + ToDelete)
		if err != nil {

			err = errors.New("数据更新失败 请稍后重试")
			return "", err
		}
		url, err := utils.SaveFile(goods.Image, c)
		if err != nil {

			tx.Rollback()
			err = errors.New("数据更新失败 请稍后重试")
			return "", err
			//err=errors.New("")
		}
		stm, err = tx.Prepare("update goods_list set url =? where Gid =?")
		if err != nil {

			err = errors.New("数据更新失败 请稍后重试")
			tx.Rollback()
			err = errors.New("数据更新失败 请稍后重试")
		}
		_, err = stm.Exec(url, goods.Gid)
		if err != nil {

			err = errors.New("数据更新失败 请稍后重试")
			tx.Rollback()
			return "", err
		}
	}
	if goods.Gname != "" {
		stm, err := tx.Prepare("update goods_list set name =? where Gid=?")
		if err != nil {

			err = errors.New("数据更新失败 请稍后重试")
			tx.Rollback()
			return "", err
		}
		_, err = stm.Exec(goods.Gname, goods.Gid)
		if err != nil {

			err = errors.New("数据更新失败 请稍后重试")
			return "", err
		}
	}
	if goods.Introduce != "" {
		stm, err := tx.Prepare("update goods_info set introduce =? where Gid=?")
		if err != nil {

			err = errors.New("数据更新失败 请稍后重试")
			tx.Rollback()
			return "", err
		}
		_, err = stm.Exec(goods.Gname, goods.Gid)
		if err != nil {

			err = errors.New("数据更新失败 请稍后重试")
			return "", err
		}
	}
	if goods.Price != 0 {
		stm, err := tx.Prepare("update goods_info set price =? where Gid=?")
		if err != nil {

			err = errors.New("数据更新失败 请稍后重试")
			tx.Rollback()
			return "", err
		}
		_, err = stm.Exec(goods.Price, goods.Gid)
		if err != nil {

			err = errors.New("数据更新失败 请稍后重试")
			return "", err
		}
	}
	err = tx.Commit()
	if err != nil {

		err = errors.New("数据更新失败 请稍后重试")
		tx.Rollback()
		return "", err
	}
	return "小商品修改成功了哦", nil
}

func DeleteGoods(gid string) (string, error) {
	tx, err := DB.Begin()
	if err != nil {

		err = errors.New("商品删除失败 再试试吧")
		tx.Rollback()
		return "", err
	}
	//var wait sync.WaitGroup

	stm, err := DB.Prepare("delete from goods_commit where Gid=?")
	if err != nil {
		err = errors.New("商品删除失败 再试试吧")
		tx.Rollback()
		return "", err
	}
	_, err = stm.Exec(gid)
	if err != nil {

		err = errors.New("商品删除失败 再试试吧")
		tx.Rollback()
		return "", err
	}
	stm, err = tx.Prepare("delete from shop_chart where gid=?")
	if err != nil {
		err = errors.New("商品删除失败 再试试吧")
		tx.Rollback()
		return "", err
	}
	_, err = stm.Exec(gid)
	if err != nil {

		err = errors.New("商品删除失败 再试试吧")
		tx.Rollback()
		return "", err
	}
	stm, err = tx.Prepare("delete from goods_info where GId=?")
	if err != nil {
		err = errors.New("商品删除失败 再试试吧")
		tx.Rollback()
		return "", err
	}
	_, err = stm.Exec(gid)
	if err != nil {

		err = errors.New("商品删除失败 再试试吧")
		tx.Rollback()
		return "", err
	}
	stm, err = DB.Prepare("select url from goods_list where Gid =?")
	var link string
	if err != nil {

		err = errors.New("商品删除失败 再试试吧")
		tx.Rollback()
		return "", err
	}
	err = stm.QueryRow(gid).Scan(&link)
	err = os.Remove("/www/static/" + link)
	if err != nil {

		err = errors.New("商品删除失败 再试试吧")
		tx.Rollback()
		return "", err
	}
	stm, err = tx.Prepare("delete  from goods_list where Gid =?")
	if err != nil {

		err = errors.New("商品删除失败 再试试吧")
		tx.Rollback()
		return "", err
	}
	_, err = stm.Exec(gid)
	if err != nil {

		err = errors.New("商品删除失败 再试试吧")
		tx.Rollback()
		return "", err
	}

	tx.Commit()
	return "商品删除成功", nil

}
