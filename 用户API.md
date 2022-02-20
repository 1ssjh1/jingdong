## 鉴权说明

### 用户 

~~~
鉴权 主要通过在head 中的 Authorization 字段中间添加  所有操作 登录时候返回的Token (无论用户或者管理着)
~~~

可能反馈信息

通用

| state | msg                      | 说明             |
| ----- | ------------------------ | ---------------- |
| false | 你是谁 验证信息不知道    | 未携带token      |
| false | 时间最宝贵 超时了        | token过期        |
| false | tmd 这根本不是我的呜呜呜 | token验证错误    |
| false | 这tm 不是token           | 非所需格式       |
| false | 你还没有登录             | cookie无或者过期 |



### 短信发送



` method :GET url :/register?Phone=xxx`

| 请求参数 | 类型 | 解释     |
| -------- | ---- | -------- |
| Phone    | 必选 | 手机号码 |

返回内容：

| 返回参数  | 解释  |
|-------|-----|
| state | 状态  |
| msg   | 信息  |


| state | msg                       | 解释 |
| ----- | ------------------------- | ---- |
| Ture  | 短信发送成功              | 无   |
| False | 短信发送失败请重试        | 无   |
| false | 短信请求失败了            |      |
| false | 短信号码错误 ，你莫呼我   |      |
| false | 手机号格式错误            |      |
| false | 手机号为空                |      |
| false | 发太快了 歇会儿再试试     |      |
| false | 今天发的太多了 明天再来吧 |      |

### 注册

`method :Post;  form 表单`

`url: /register `

| 请求参数 | 类型 | 解释   |
| --------- | ---- | ------ |
| username  | 必选 | 用户名 |
| password  | 必选 | 密码   |
| code      | 必选 | 验证码 |
| number    | 必选 | 验证码 |

返回参数

 |返回参数| 解释|
|---|---|
|state|状态|
|msg|信息|


返回内容

| status | msg                | 解释           |
|--------|--------------------| -------------- |
| false  | 参数绑定失败             | 参数错误       |
| false  | 验证码超时              | 验证码超时 |
| false  | 验证码错误              |验证码错误|
| false  | 用户名已存在 要不登录试试      | 重复注册       |
| false  | 注册失败               | 数据库操作失败 |
| false  | 手机号已经被注册           | 手机号重复绑定 |
| ture   | 注册成功               | 可以Happy了    |

### 登录

`Post :/login ;  form-date`

请求内容

| 请求参数 | 类型 | 说明     |
| -------- | ---- | -------- |
| username | 必选 | 用户名   |
| password | 必选 | 用户密码 |

返回参数

|返回参数|解释|
|---|---|
|state|状态|
|msg|解释|
|Token|认证Token|

| state | msg                | 说明              |
| ----- | ------------------ | ----------------- |
| false | 参数绑定失败       | 缺少参数          |
| false | 登录失败           | 数据库操作失败    |
| false | 你还没注册登录个屁 | 用户未注册        |
| false | 密码错误           | 无                |
| true  | " "         | 可以开心的Happy了 |


 ### 退出登录

`method Get url /logout`

| 返回参数| 解释   |
|---|------|
|state| 状态   |
|msg| 信息解释 |

返回内容

|state|msg|解释|
|---|---|---|
|false|退出登录失败|未知错误|
|false|你丫还没登录呢|用户未登录|
|true|退出登录成功|无|

### 新闻接口

`url :/news METHOD :post    `

返回参数

| 返回参数 | 说明     |
| -------- | -------- |
| state    | 状态     |
| msg      | 新闻信息 |

返回示例

~~~
{
  "status": 200,
  "data": [
    {
      "articleid": 10224,
      "title": "2021----2022学年第一学期结换毕学生成绩公布"
    },
    {
      "articleid": 10220,
      "title": "重庆邮电大学“2021级计算机科学文峰实验班”招生简章"
    },
    {
      "articleid": 10204,
      "title": "转发物业管理中心关于2021-2022年第二学期教材预定的通知"
    },
    {
      "articleid": 10126,
      "title": "2021-2022学年第一学期考试公告合集"
    }
  ]
}
~~~





## 用户模块
 ### 查询余额
`method Get url /user/balance form-date`



|返回参数| 说明         |
|---|------------|
|state| 状态         |
|msg| 信息         |
|balance| 余额（成功才会返回） |

返回内容

|state|msg| balance              |
|---|---|----------------------|
|false|参数绑定失败| 缺少参数                 |
|false|身份验证失败| 身份验证错误               |
|false|查找失败| 信息查找失败               |
|false|未查到用户信息| 信息缺失（这里会改 先检验用户是否存在） |

### 充值余额
` method post url /user/balance form date`

请求内容


|请求参数|类型|说明|
|---|---|---|
|balance|必选|充值金额|



|返回参数|说明|
|---|---|
|state|状态|
|msg|信息|


返回内容

|state|msg|说明|
|---|---|---|
|false|参数绑定失败|参数错误|
|false|用户信息不匹配|身份校验失败|
|false|充值失败|数据库操作失败|
|true|充值成功|

### 查看订单

`method POST url:/user/order form-date`

返回参数

| 返回参数 | 说明                          |
| -------- | ----------------------------- |
| state    | 状态                          |
| msg      | 失败返回原因/成功返回所有信息 |

返回实例

| state | msg          |
| ----- | ------------ |
| false | 参数错误     |
| false | 身份验证失败 |
| false | nil          |
| true  | 略           |

### 确认收货

`method PUT url:/user/order form-date`

请求参数

| 请求参数 | 说明   |
| -------- | ------ |
| oid      | 订单id |

返回参数

| 返回参数 | 说明     |
| -------- | -------- |
| state    | 状态     |
| msg      | 详细信息 |

返回实例

| state | msg               | 说明           |
| ----- | ----------------- | -------------- |
| false | 参数绑定失败      |                |
| false | 身份验证失败      |                |
| false | 失败              | 数据库操作失败 |
| false | 还没发货呢 着啥急 |                |
| false | 错误              |                |
| true  | 确认收货成功      |                |

### 删除订单

`mehthod DELETE url:/user/order  form-date`

请求参数

| 请求参数 | 说明   |
| -------- | ------ |
| oid      | 订单id |

返回参数

| 返回参数 | 说明 |
| -------- | ---- |
| state    | 状态 |
| msg      | 信息 |

返回实例

| state | msg                       |
| ----- | ------------------------- |
| false | 参数绑定失败              |
| false | 身份验证失败              |
| false | 失败                      |
| false | 怎么可以动别人的订单呢    |
| false | 已经完成订单 概不负责了哦 |
| true  | 订单销毁成功              |



## 商城模块

### 获取所有商品信息

`method POST url: /shop/all  form -date` 




|返回参数| 说明                        |
|---|---------------------------|
|state| 状态                        |
|msg| 信息(请求正确时返回所有商品信息，失败返回失败原因）|

msg 详解

``` 

[ 
  一条商品信息
    {
        "Gid": 1,
        "Name": "快乐",
        "Sales": 100,
        "Commit": 1,
        "Grate": 100,
        "Introduce": "售卖快乐",
        "Choose": [
        
        //Choose 为商品种类
        
            {
                "Cid": 1,
                "Types": "轻度快乐",
                "Price": 100
            },
            {
                "Cid": 2,
                "Types": "重度快乐",
                "Price": 120
            }
        ]
    },

]
```

### 加入购物车

` method :post   Url: /shop/chart form-date`

请求内容

| 参数  | 类型 | 说明   |
| :---: | ---- | ------ |
|  gid  | 必选 | 商品id |
| count | 必选 | 数量   |

返回参数

| 返回参数 | 说明 |
| -------- | ---- |
| msg      | 信息 |
| state    | 状态 |

返回实例

| state | msg                          | 说明           |
| ----- | ---------------------------- | -------------- |
| false | 参数绑定错误                 | 缺少参数       |
| false | 身份验证失败                 | 身份校验失败   |
| false | 已经加入购物车 试试别的吧    | 重复加入       |
| false | ""                           | 数据库操作失败 |
| true  | 你的宝贝已经躺在购物车里了哦 |                |

### 获取购物车信息

`method GET url:/shop/chart  form-date`  

请求参数

返回参数

| 返回参数 | 说明                                        |
| -------- | ------------------------------------------- |
| state    | 状态                                        |
| msg      | 信息（成功返回所有购物车信息 失败返回原因） |

返回实例

| state | msg            |
| ----- | -------------- |
| false | 参数绑定错误   |
| false | 查询失败       |
| false | 身份验证失败   |
| true  | 所有购物车信息 |





### 修改购物车

` method PUT url:/shop/chart form-date`



请求参数

` count ==0 表示从购物车移除该商品 可供修改cid 和 count  `

| 请求参数 | 类型 | 说明   |
| -------- | ---- | ------ |
| gid      | 必选 | 商品id |
| cid      | 必选 | 类型Id |
| count    | 必选 | 数量   |

返回参数



| 返回参数 | 说明     |
| -------- | -------- |
| state    | 状态     |
| msg      | 详细信息 |

返回实例



| state | msg                    |
| ----- | ---------------------- |
| false | 参数绑定失败           |
| false | 身份验证失败           |
| false | 失败（数据库操作失败） |
| false | 用户信息不匹配         |
| true  | 宝贝忍痛离开了购物车   |
| true  | 操作失败               |

### 生成订单

` method POST url:shop/order form-date`

请求参数

| 请求参数 | 说明                    |
| -------- | ----------------------- |
| chart_id | 购物车商品id 可提交多个 |

返回参数

| 返回参数 | 说明     |
| -------- | -------- |
| state    | 状态     |
| msg      | 具体信息 |

返回实例

| state | msg                    |
| ----- | ---------------------- |
| false | 参数绑定失败           |
| false | 身份验证失败           |
| false | 错误（数据库操作失败） |
| false | 订单创建失败           |
| false | 你个穷逼               |
| true  | 订单创建成功           |



### 轮播图

method :GET url 

| 请求参数 | 类型 | 解释 |
| -------- | ---- | ---- |
| id       | 必选 | 无   |

| 返回参数 | 说明   |
| -------- | ------ |
| 状态     | 状态   |
| msgsrc   | msgsrc |




msgsrc详解：

```
{
lunbo001:{
        "src001":"",
        "src002":"",
        "src003":"",
        "src004":"",
        "src005":""
       {    
}


{
lunbo002:{
        "src001":"",
        "src002":"",
        "src003":"",
        "src004":"",
        "src005":""，
        "src006":"",
        "src007":"",
        "src008":"",
        "src009":""  
       {    
}


```

### 用户主界面

` url :/user/    method :  GET     form-date`



返回参数 

| 返回参数 | 说明                                       |
| -------- | ------------------------------------------ |
| state    | 状态                                       |
| msg      | 信息 （无误返回用户信息 错误返回错误原因） |

失败实例

| state | msg              |
| ----- | ---------------- |
| false | 参数绑定失败     |
| false | 个人信息读取失败 |
|       |                  |



成功实例


```go
{
  "msg": {
    "Uid": 1,
    "Username": "Siana",
    "Balance": 25034,
    "ImageUrl": "hfg",
    "Category": [
      {
        "State": "已支付",
        "Order": [
          {
            "Gid": 2,
            "Oid": 27,
            "Count": 56,
            "State": "已支付"
          },
          {
            "Gid": 1,
            "Oid": 28,
            "Count": 56,
            "State": "已支付"
          },
          {
            "Gid": 1,
            "Oid": 30,
            "Count": 2,
            "State": "已支付"
          }
        ]
      },
      {
        "State": "已发货",
        "Order": [
          {
            "Gid": 1,
            "Oid": 29,
            "Count": 56,
            "State": "已发货"
          }
        ]
      },
      {
        "State": "已完成",
        "Order": [
          {
            "Gid": 2,
            "Oid": 26,
            "Count": 56,
            "State": "已完成"
          }
        ]
      }
    ]
  },
  "state": true
}
```

### 找回密码

` url :/find   methed POST  form-date` 

请求参数 

| 请求参数 | 说明                                      |
| -------- | ----------------------------------------- |
| username | 用户名                                    |
| password | 修改的密码                                |
| code     | 验证码（发送表单前 使用短信接口发送短信） |
| number   | 手机号                                    |

返回参数

| 返回参数 | 说明 |
| -------- | ---- |
| state    | 状态 |
| msg      | 信息 |

返回实例

| state | msg                            | 说明 |
| ----- | ------------------------------ | ---- |
| false | 参数绑定失败                   |      |
| false | 验证码超时                     |      |
| false | 验证码错误                     |      |
| false | 账户找回失败（数据库操坐失败） |      |
| false | 手机号不匹配，再试试吧         |      |
| false | 密码不能和原来的相同哦         |      |
| true  |                                |      |

### 首页展示

~~~
url : /show  method :Get 
// 无请求参数
返回示例
{
    "msg": {
        "All": [
            {
                "Gid": 1,
                "Name": "快乐",
                "Url": "https://sanser.ltd/static/f's'd",
                "Type": "ghfd",
                "Price": 100,
                "Sales": 100,
                "Commit": 1,
                "Grate": 100,
                "Introduce": "售卖快乐"
            },
            {
                "Gid": 2,
                "Name": "快乐水",
                "Url": "https://sanser.ltd/static/t're",
                "Type": "hgf",
                "Price": 200,
                "Sales": 1200,
                "Commit": 100,
                "Grate": 95,
                "Introduce": "快乐水"
            },
            {
                "Gid": 3,
                "Name": "fsadfa",
                "Url": "https://sanser.ltd/static/654",
                "Type": "hgf",
                "Price": 534,
                "Sales": 52,
                "Commit": 52,
                "Grate": 52,
                "Introduce": "股市大幅改善"
            },
            {
                "Gid": 4,
                "Name": "rog全家桶 一套",
                "Url": "https://sanser.ltd/static/1645017515.webp",
                "Type": "京东秒杀",
                "Price": 8888,
                "Sales": 0,
                "Commit": 0,
                "Grate": 100,
                "Introduce": "ROG幻16 16英寸设计师高性能游戏笔记本电脑(i7-12700H"
            },
            {
                "Gid": 5,
                "Name": "施华洛世奇项链",
                "Url": "https://sanser.ltd/static/1645019976.webp",
                "Type": "京东秒杀",
                "Price": 520,
                "Sales": 0,
                "Commit": 0,
                "Grate": 100,
                "Introduce": "SPARKLING DC项链情人节礼物送女友"
            },
            {
                "Gid": 6,
                "Name": "RedRock 5G手机",
                "Url": "https://sanser.ltd/static/1645020051.webp",
                "Type": "京东秒杀",
                "Price": 1,
                "Sales": 0,
                "Commit": 0,
                "Grate": 100,
                "Introduce": "22.5W超级快充 6nm 5G疾速芯 全网通版 6G+128G"
            },
            {
                "Gid": 7,
                "Name": "表情包0",
                "Url": "https://sanser.ltd/static/1645020169.jpg",
                "Type": "精选部分",
                "Price": 888,
                "Sales": 0,
                "Commit": 0,
                "Grate": 100,
                "Introduce": "象征redrock的压迫感"
            },
            {
                "Gid": 8,
                "Name": "表情包1",
                "Url": "https://sanser.ltd/static/1645020204.jpg",
                "Type": "精选部分",
                "Price": 888,
                "Sales": 0,
                "Commit": 0,
                "Grate": 100,
                "Introduce": "象征达超导师的帅气"
            },
            {
                "Gid": 9,
                "Name": "表情包2",
                "Url": "https://sanser.ltd/static/1645020228.jpg",
                "Type": "精选部分",
                "Price": 888,
                "Sales": 0,
                "Commit": 0,
                "Grate": 100,
                "Introduce": "象征兴兴学姐的温柔"
            },
            {
                "Gid": 10,
                "Name": "表情包3",
                "Url": "https://sanser.ltd/static/1645020247.jpg",
                "Type": "精选部分",
                "Price": 888,
                "Sales": 0,
                "Commit": 0,
                "Grate": 100,
                "Introduce": "象征文轩导师的负责"
            },
            {
                "Gid": 11,
                "Name": "表情包4",
                "Url": "https://sanser.ltd/static/1645020267.jpg",
                "Type": "精选部分",
                "Price": 888,
                "Sales": 0,
                "Commit": 0,
                "Grate": 100,
                "Introduce": "象征闻捷姐姐的温柔"
            },
            {
                "Gid": 12,
                "Name": "表情包5",
                "Url": "https://sanser.ltd/static/1645020287.jpg",
                "Type": "精选部分",
                "Price": 888,
                "Sales": 0,
                "Commit": 0,
                "Grate": 100,
                "Introduce": "象征佳俊导师的认真"
            },
            {
                "Gid": 13,
                "Name": "可乐",
                "Url": "https://sanser.ltd/static/1645020368.jpg",
                "Type": "美食部分",
                "Price": 5490,
                "Sales": 0,
                "Commit": 0,
                "Grate": 100,
                "Introduce": "肥仔快乐"
            },
            {
                "Gid": 14,
                "Name": "可口小番茄",
                "Url": "https://sanser.ltd/static/1645020406.webp",
                "Type": "美食部分",
                "Price": 77,
                "Sales": 0,
                "Commit": 0,
                "Grate": 100,
                "Introduce": "健康轻食千禧西红柿小番茄口感圣女果水"
            },
            {
                "Gid": 15,
                "Name": "机械键盘",
                "Url": "https://sanser.ltd/static/1645020456.webp",
                "Type": "百货部分",
                "Price": 678,
                "Sales": 0,
                "Commit": 0,
                "Grate": 100,
                "Introduce": "粉色少女心，满足你的幻想"
            },
            {
                "Gid": 16,
                "Name": "UFO",
                "Url": "https://sanser.ltd/static/1645020476.webp",
                "Type": "百货部分",
                "Price": 77,
                "Sales": 0,
                "Commit": 0,
                "Grate": 100,
                "Introduce": "UFO魔术悬浮球反重力黑科技魔法棒"
            },
            {
                "Gid": 17,
                "Name": "YSL精选口红",
                "Url": "https://sanser.ltd/static/1645020524.webp",
                "Type": "个护部分",
                "Price": 999,
                "Sales": 0,
                "Commit": 0,
                "Grate": 100,
                "Introduce": "YSL口红全新烈艳蓝金唇膏哑光"
            },
            {
                "Gid": 18,
                "Name": "fufu",
                "Url": "https://sanser.ltd/static/1645020558.png",
                "Type": "个护部分",
                "Price": 888,
                "Sales": 0,
                "Commit": 0,
                "Grate": 100,
                "Introduce": "fufu 初音未来玩偶 哔哩哔哩 毛绒玩偶公仔娃娃靠枕国产版 公服式 32厘米"
            }
        ]
    },
    "state": true
}
~~~



### 用户头像更新



` url : /user/image method :PUT form-date` 

| 请求参数 | 说明            |
| -------- | --------------- |
| image    | 头像文件 5M以内 |

返回参数

| 返回参数 | 说明     |
| -------- | -------- |
| state    | 状态     |
| msg      | 反馈信息 |
|          |          |

返回实例

| state | msg                   |
| ----- | --------------------- |
| false | 参数绑定失败          |
| false | 参数缺失              |
| false | 文件过大 换个小点的吧 |
| false | 文件保存失败          |
| true  | 文件上传成功          |

### 提交评论

` url  :/user/commit method POST form-date`

请求参数

| 请求参数 | 说明   |
| -------- | ------ |
| oid      | 订单id |
| commit   | 评论   |
|          |        |

返回参数

| 返回参数 | 说明     |
| -------- | -------- |
| state    | 状态     |
| msg      | 反馈信息 |

返回实例

| state | msg          |
| ----- | ------------ |
| false | 参数绑定失败 |
| false | 参数缺失     |
| false | 评论提交失败 |
| false | 商品状态错误 |
| true  | 评论提交成功 |

### 获取评论

~~~go
url :/shop/commit?gid=? method:GET   
~~~



## 管理系统接口

###  登录

`url :/admin/login method :POST form-date`

请求参数

| 请求参数 | 说明   |
| -------- | ------ |
| username | 用户名 |
| password | 密码   |

返回参数

| 返回参数 | 说明                             |
| -------- | -------------------------------- |
| state    | 状态                             |
| msg      | 附带信息                         |
| token    | 验证token 登录成功才会返回此信息 |

返回示例

| state | msg          | token     |
| ----- | ------------ | --------- |
| false | 参数绑定失败 |           |
| false | 登录失败     |           |
| false | 密码错误     |           |
| true  | 登陆成功     | 附带token |

### 订单查看

`url :/admin/order method :GET head 附带token`

返回参数

| 返回参数 | 说明 |
| -------- | ---- |
| state    | 状态 |
| msg      | 信息 |

返回实例

| 返回参数 | 说明         |
| -------- | ------------ |
| false    | 参数缺失     |
| false    | 数据查询失败 |
| true     | 所有信息     |

### 更新订单

`url :/admin/order method :PUT form-date`

请求参数

| 说明           |
| -------------- |
| 订单id         |
| 想要修改的状态 |



返回参数

| 返回参数 | 说明     |
| -------- | -------- |
| state    | 状态     |
| msg      | 附带信息 |

返回实例

| state | msg          |
| ----- | ------------ |
| false | 用户信息错误 |
| false | 订单不存在   |
| false | 操作失败     |
| true  | 操作成功     |

### 删除订单

`url /admin/order method:delete form-date`

请求参数

| 请求参数 | 说明     |
| -------- | -------- |
| oid      | 操作的id |

返回参数

| 返回参数 | 说明 |
| -------- | ---- |
| state    | 状态 |
| msg      | 信息 |

返回实例

| state | msg          |
| ----- | ------------ |
| false | 订单不存在   |
| false | 操作失败     |
| false | 用户信息错误 |
| true  | 操作成功     |

### 添加商品

`url:/admin/goods method :POST form -date`

请求参数

| 请求参数  | 说明     |
| :-------: | -------- |
|   image   | 图像文件 |
|   gname   | 商品名称 |
|   type    | 分类     |
| introduce | 介绍     |
|   price   | 定价     |

返回参数

| 返回参数 | 说明 |
| -------- | ---- |
| state    | 状态 |
| msg      | 信息 |

返回实例

| state | msg                       |
| ----- | ------------------------- |
| false | 商品添加失败 请重试       |
| false | 文件过大 换个小点的试试吧 |
| true  | 商品添加成功              |



### 获取商品信息 

~~~
url :/admin/goods  GET 参数 响应 同/shop/all
~~~

### 更新商品信息

请求参数

| 请求参数  | 说明                                              |
| :-------: | ------------------------------------------------- |
|    gid    | 操作的商品id// 以下参数 可传可不传 传输则代表修改 |
|   image   | 图像文件                                          |
|   gname   | 商品名称                                          |
|   type    | 分类                                              |
| introduce | 介绍                                              |
|   price   | 定价                                              |

### 删除商品信息

`url :/admin/goods method :DELETE form-date`

| 请求参数 | 说明   |
| -------- | ------ |
| Gid      | 商品id |

| 返回参数 | 说明 |
| -------- | ---- |
| state    | 状态 |
| msg      | 信息 |

返回实例、

| state | msg                   |
| ----- | --------------------- |
| false | 商品删除失败 再试试吧 |
| true  | 商品删除成功          |

### 登出

~~~、
url:/admin/logout method :get
~~~

| 返回参数 | 说明 |
| -------- | ---- |
| state    | 状态 |
| msg      | 信息 |

| state | msg          |
| ----- | ------------ |
| true  | 退出登录成功 |



