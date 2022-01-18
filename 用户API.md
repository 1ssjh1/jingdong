







### 用户注册

请求方式：“post”

| 请请求参数 | 类型 | 解释             |
| ---------- | ---- | ---------------- |
| phone      | 必选 | 手机号码         |
| password   | 必选 | 密码             |
| nicknak    | 可选 | 昵称             |
| capcha     | 可选 | 验证码（加分项） |





| status | data               | 解释               |
| ------ | ------------------ | ------------------ |
| false  | “用户名不能为空”   | nickname为空       |
| false  | “用户名过长”       | nickname超过12字节 |
| false  | "用户名过短"       | nickname小于6字节  |
| false  | “密码过长”         | password超过12字节 |
| false  | “密码过短”         | password小于6字节  |
| false  | “手机号已经被注册” | phone已经被注册    |
| false  | “手机号为空”       | phone为空          |
| ture   | “1”                | 参数正确           |

