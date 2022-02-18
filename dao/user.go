package dao

import (
	"JD/models"
	"errors"
	"fmt"
)

func Register(name string, word string, number string) error {
	var U models.Register
	stm, err := DB.Prepare("select  name from user_info where name = ?")
	if err != nil {
		fmt.Println(err)
		err = errors.New("注册失败")
		return err
	}
	defer stm.Close()
	rows, err := stm.Query(&name)
	if err != nil {
		fmt.Println(err)
		err = errors.New("注册失败")
		return err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&U.Username)
		if name == U.Username {
			err = errors.New("用户名已存在 要不登录试试")
			return err
		}
	}
	stm, err = DB.Prepare("select  number from user_info where number = ?")
	if err != nil {
		fmt.Println(err)
		err = errors.New("注册失败")
		return err
	}
	defer stm.Close()
	rows, err = stm.Query(&number)
	if err != nil {
		fmt.Println(err)
		err = errors.New("注册失败")
		return err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&U.Username)
		if number == U.Username {
			err = errors.New("手机号已注册")
			return err
		}
	}
	stm, err = DB.Prepare("insert into user_info(name,word,number) values(?,?,?);")
	if err != nil {
		fmt.Println(err)
		err = errors.New("注册失败")
		return err
	}
	_, err = stm.Exec(name, word, number)
	if err != nil {
		fmt.Println(err)
		err = errors.New("注册失败")
		return err
	}
	return nil

}
func Find(user models.Register) error {
	stm, err := DB.Prepare("select word ,number from user_info where name=?")
	if err != nil {
		err = errors.New("账户找回失败")
		return err
	}
	var temple models.Register
	err = stm.QueryRow(user.Username).Scan(&temple.Password, &temple.Number)
	if err != nil {
		err = errors.New("账户找回失败")
		return err
	}
	if temple.Number != user.Number {
		err = errors.New("手机号不匹配，再试试把")
		return err
	}
	if temple.Password == user.Password {
		err = errors.New("密码不能和原来的相同哦")
		return err
	}
	stm, err = DB.Prepare("update user_info set word =? where name =?")
	if err != nil {
		err = errors.New("账号找回失败")
	}
	_, err = stm.Exec(user.Password, user.Username)
	if err != nil {
		err = errors.New("账号找回失败")
		return err
	}
	return nil

}
func Login(u models.Login) (*models.BasicInfo, error) {
	stm, err := DB.Prepare("select uid, word from user_info where name = ?")
	if err != nil {
		fmt.Println(err)
		err = errors.New("登录失败")
		return nil, err
	}
	defer stm.Close()
	var basicinfo models.BasicInfo
	rows, err := stm.Query(u.Username)
	if err != nil {
		fmt.Println(err)
		err = errors.New("登录失败")
		return nil, err
	}
	var tmp models.Login
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&basicinfo.Uid, &tmp.Password)
	}
	if tmp.Password == "" {
		err = errors.New("你还没注册登录个屁")
		return nil, err
	}
	if tmp.Password != u.Password {
		err = errors.New("密码错误")
		return nil, err

	}
	basicinfo.Username = u.Username
	return &basicinfo, nil
}
func AdminLogin(name string, word string) (bool, string) {
	stm, err := DB.Prepare("select  word from admin_info where name = ?")
	if err != nil {
		return false, "登录失败"
	}
	rows, err := stm.Query(name)
	if err != nil {
		return false, "登录失败"
	}
	var s string
	for rows.Next() {
		rows.Scan(&s)
	}
	if s != word {
		return false, "密码错误"
	}
	return true, "密码正确"

}

func SaveFile(url string, user models.BasicInfo) (string, error) {
	stm, err := DB.Prepare("update user_info set image=? where uid=? ")
	if err != nil {
		err := errors.New("文件上传失败")
		return "", err
	}
	_, err = stm.Exec(url, user.Uid)
	if err != nil {
		err := errors.New("文件上传失败")
		return "", err
	}
	return "文件上传成功", nil
}

func GetBalance(u string) (bool, interface{}) {
	stm, err := DB.Prepare("select  balance from user_info where name = ?")
	if err != nil {
		return false, "查找失败"
	}
	defer stm.Close()
	rows, err := stm.Query(u)
	if err != nil {
		return false, "查找失败"
	}
	var tmp int
	for rows.Next() {
		rows.Scan(&tmp)
	}

	return true, tmp

}
func ChargeBalance(u models.Balance) (bool, string) {
	stm, err := DB.Prepare("select balance from user_info where uid=?")
	if err != nil {
		return false, "充值失败"
	}
	var balance int
	err = stm.QueryRow(u.Uid).Scan(&balance)
	balance += u.Balance
	stm, err = DB.Prepare("update user_info set balance=? where name=?")
	if err != nil {
		return false, "充值失败"
	}
	_, err = stm.Exec(balance, u.Username)
	if err != nil {
		return false, "充值失败"
	}
	return true, "充值成功"
}
func AddChart(chart models.AddChart) (bool, string) {
	stm, err := DB.Prepare("select gid from shop_chart where uid=?")
	if err != nil {
		fmt.Println(err)
		return false, ""
	}
	var Temple models.AddChart
	rows, err := stm.Query(chart.Uid)
	for rows.Next() {
		rows.Scan(&Temple.Gid)
		if Temple.Gid == chart.Gid {
			return false, "已经加入购物车了 试试别的吧"
		}
	}

	stm, err = DB.Prepare(" insert into shop_chart(uid,gid,count) values(?,?,?)")
	if err != nil {
		fmt.Println(err, "fod")

		return false, ""
	}
	_, err = stm.Exec(chart.Uid, chart.Gid, chart.Count)
	if err != nil {
		fmt.Println(err)

		return false, ""
	}
	return true, "你的宝贝已经躺在购物车里了哦"

}
func AllChart(user models.Userinfo) (*models.AllChart, error) {
	stm, err := DB.Prepare("select uid from user_info where name=?")
	if err != nil {
		fmt.Println(err)
		err = errors.New("查询失败")
		return nil, err
	}
	var Temple models.Userinfo
	err = stm.QueryRow(user.Username).Scan(&Temple.Uid)
	if err != nil {
		fmt.Println(err)
		err = errors.New("查询失败")
		return nil, err
	}
	if Temple.Uid != user.Uid {
		fmt.Println("参数错误")
		err = errors.New("参数错误")
		return nil, err
	}

	all := models.AllChart{}
	all.BasicInfo = user.BasicInfo
	stm, err = DB.Prepare("select chart_id,gid,count from shop_chart where uid=? ")
	if err != nil {
		fmt.Println(err)
		err = errors.New("查询失败")
		return nil, err
	}
	rows, err := stm.Query(all.Uid)
	if err != nil {
		fmt.Println(err)
		err = errors.New("查询失败")
		return nil, err
	}
	Templevalue := models.ChartShop{}
	for rows.Next() {
		rows.Scan(&Templevalue.ChartId, &Templevalue.Gid, &Templevalue.Count)
		stm, err = DB.Prepare("select name from goods_list where Gid=? ")
		err = stm.QueryRow(Templevalue.Gid).Scan(&Templevalue.Good)
		all.ChartList = append(all.ChartList, Templevalue)
	}
	return &all, nil

}
func UpdateChart(chart models.ShopChart) (bool, string) {

	//将Count设置为0 就意味着删除
	if chart.Count == 0 {

		//查看数据是否合法
		stm, err := DB.Prepare("select uid  from shop_chart where chart_id=? ")
		var Temple models.ShopChart
		err = stm.QueryRow(chart.ChartId).Scan(&Temple.Uid)
		if err != nil {
			fmt.Println(err)
			return false, "失败"
		}
		if Temple.Uid != chart.Uid {
			return false, "你怎么能动别人的订单呢"
		}

		stm, err = DB.Prepare("delete from shop_chart where chart_id=?")
		if err != nil {
			return false, "失败"

		}
		_, err = stm.Exec(chart.ChartId)
		if err != nil {
			return false, "失败"

		}
		return true, "宝贝忍痛离开了购物车"

	}
	stm, err := DB.Prepare("update shop_chart set count =?  where chart_id=? ")
	if err != nil {
		fmt.Println(err)
		return false, "失败"
	}
	_, err = stm.Exec(chart.Count, chart.ChartId)
	if err != nil {
		return false, "失败"
	}
	return true, "操作成功"
}

func MyInfo(User models.BasicInfo) (*models.MyInfo, error) {
	var UserInfo models.MyInfo
	//fmt.Println(User, "userinfomation")
	UserInfo.BasicInfo = User

	fmt.Println(UserInfo.Uid)
	stm, err := DB.Prepare("select balance ,image from user_info where uid=?")
	if err != nil {
		err = errors.New("个人信息读取错误")
		return nil, err
	}
	fmt.Println(UserInfo)
	var tmp models.MyInfo
	rows, err := stm.Query(User.Uid)
	err = stm.QueryRow(User.Uid).Scan(&UserInfo.Balance, &UserInfo.ImageUrl)
	fmt.Println(err)
	fmt.Println(UserInfo.ImageUrl)
	if err != nil {
		err = errors.New("个人信息读取错误")
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&tmp.Balance, &tmp.ImageUrl)
		fmt.Println(tmp)
		if err != nil {
			err = errors.New("个人信息读取错误")
			return nil, err
		}
		UserInfo = tmp
	}
	var link = "https://sanser.ltd/static/"

	UserInfo.ImageUrl = link + tmp.ImageUrl
	stm, err = DB.Prepare("select State from Order_state_type ")
	if err != nil {
		err = errors.New("个人信息读取错误")
		return nil, err
	}
	rows, err = stm.Query()
	if err != nil {
		err = errors.New("个人信息读取错误")
		return nil, err
	}
	var Category models.Category
	for rows.Next() {
		err = rows.Scan(&Category.State)
		if err != nil {
			err = errors.New("个人信息读取错误")
			return nil, err
		}
		UserInfo.Category = append(UserInfo.Category, Category)
	}
	stm, err = DB.Prepare("select gid ,oid ,state,count from user_order where uid =? and state=? ")
	var temple models.AllOrder
	for key, value := range UserInfo.Category {
		rows, err = stm.Query(User.Uid, value.State)
		if err != nil {
			err = errors.New("个人信息读取错误")
			return nil, err
		}
		for rows.Next() {
			err = rows.Scan(&temple.Gid, &temple.Oid, &temple.State, &temple.Count)
			if err != nil {
				err = errors.New("个人信息读取错误")
				return nil, err
			}
			UserInfo.Category[key].Order = append(UserInfo.Category[key].Order, temple)
		}
	}
	UserInfo.BasicInfo = User

	return &UserInfo, nil
}
