# 管理员API文档

## 管理员仪表板
请求方法: GET

URL: localhost:8080/admin/dashboard

描述: 获取管理员仪表板或主页的信息。

## 查看所有职位发布

请求方法: GET

URL: localhost:8080/admin/jobs

描述: 查看系统中所有职位发布的列表。

## 查看职位详情

请求方法: GET

URL: localhost:8080/admin/job/:jobID

描述: 根据职位ID查看特定职位的详细信息。

| 参数名称 | 必选 | 类型   | 描述         |
|----------|------|--------|--------------|
| academy  | 是   | String | 用户所在学院 |
| major    | 是   | String | 用户的专业   |
| id       | 是   | String | 用户ID       |
| name     | 是   | String | 用户姓名     |
| job      | 是   | String | 用户职业     |
| state    | 是   | String | 用户状态     |


## 查看职位申请

请求方法: GET

URL: localhost:8080/admin/applications/:jobID

描述: 查看特定职位的申请人列表。

### 输出参数
| 参数名称 | 必选 | 类型   | 描述         |
|----------|------|--------|--------------|
| academy  | 是   | String | 用户所在学院 |
| major    | 是   | String | 用户的专业   |
| id       | 是   | String | 用户ID       |
| name     | 是   | String | 用户姓名     |
| state    | 是   | String | 用户状态     |

## 查看所有职位申请

请求方法: GET

URL: localhost:8080/admin/applications/all

描述: 查看系统中所有职位的申请。

### 输出参数

| 参数名称 | 必选 | 类型   | 描述         |
|----------|------|--------|--------------|
| academy  | 是   | String | 用户所在学院 |
| major    | 是   | String | 用户的专业   |
| id       | 是   | String | 用户ID       |
| name     | 是   | String | 用户姓名     |
| job      | 是   | String | 用户职业     |
| state    | 是   | String | 用户状态     |

## 获取未处理的职位申请

请求方法: GET

URL: localhost:8080/admin/applications/unhandled

描述: 获取所有未处理的职位申请。

### 输出参数

| 参数名称 | 必选 | 类型   | 描述         |
|----------|------|--------|--------------|
| academy  | 是   | String | 用户所在学院 |
| major    | 是   | String | 用户的专业   |
| id       | 是   | String | 用户ID       |
| name     | 是   | String | 用户姓名     |

## 获取特定职位的未处理申请
请求方法: GET

URL: localhost:8080/admin/applications/unhandled/:jobID

描述: 获取特定职位的未处理申请。

### 输出参数

| 参数名称 | 必选 | 类型   | 描述         |
|----------|------|--------|--------------|
| academy  | 是   | String | 用户所在学院 |
| major    | 是   | String | 用户的专业   |
| id       | 是   | String | 用户ID       |
| name     | 是   | String | 用户姓名     |
| job      | 是   | String | 用户职业     |

## 审批或拒绝职位申请

请求方法: PUT

URL: localhost:8080/admin/application/approve

描述: 对职位申请进行审批或拒绝。


