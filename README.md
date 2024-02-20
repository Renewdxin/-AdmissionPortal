## 地址

git仓库地址：先提交到各自的分支

后端：back_branch

前端：front_branch

```
git@github.com:Renewdxin/-AdmissionPortal.git
```

## 后端ip地址

```
127.0.0.1:8080
//后期服务器
unknown
```

## 路由
```go
    //home page
    r.POST("/home", homeHandler.HomePage)
    
    // auth setting
	apiAccount := r.Group("/auth")
	apiAccount.Use()
	{
        apiAccount.POST("/login", authHandler.Login)
        //apiAccount.POST("/logout")
        apiAccount.POST("/signup", authHandler.Register)
        apiAccount.POST("/code/send", authHandler.CodeSend)
        apiAccount.POST("/code/verify", authHandler.CodeVerify)
        apiAccount.POST("/password/change/code", authHandler.ChangePasswordByCode)
        apiAccount.POST("/password/change/pwd", authHandler.ChangePasswordByPwd)
	}
    
    // personal info
    apiProfile := r.Group("/profile")
    apiProfile.Use(jwtHandler.JWTHandler())
    {
        // GetUserInfo 通过id得到用户信息，返回项为姓名、性别、出生日期、邮箱、手机号，空则为""
        apiProfile.GET("/Info/:id", userHandler.GetUserInfo)
        // DeleteUser 手机验证码/密码 验证才能删除
        apiProfile.DELETE("/delete/:id", userHandler.DeleteUser)
        // 更新用户信息，仅限手机号、邮箱
        apiProfile.PUT("/update/:id", userHandler.UpdateUserInfo)
        // 查询是否被录取
        apiProfile.GET("/status/:id", userHandler.GetUserStatus)
    }
    
    apiJob := r.Group("/recruitment")
    apiJob.Use()
    {
        // 查看岗位总览
        apiJob.GET("/jobs", jobHandler.GetJobs)
        // 查看岗位详细信息
        apiJob.GET("/job/:jobID", jobHandler.GetJobInfo)
        // 用户申请投递
        apiJob.POST("/job/:jobID/apply/:userID", jobHandler.ApplyJob)
    }
    
    apiAdmin := r.Group("/admin")
    apiAdmin.Use() // 使用JWT中间件进行管理员身份验证
    {
        // 管理员仪表板或主页
        apiAdmin.GET("/dashboard", adminHandler.HomePage)
        // 查看所有职位发布（管理员）
        apiAdmin.GET("/jobs", adminHandler.ShowAllJobs)
        // 查看职位详情（管理员）
        apiAdmin.GET("/job/:jobID", adminHandler.ShowJobDetails)
        // 查看职位申请（管理员）
        apiAdmin.GET("/applications/:jobID", adminHandler.ShowJobApply)
        // 审批或拒绝职位申请（管理员）+ 自动发送录取短信
        apiAdmin.PUT("/application/:appID", adminHandler.ApproveJobs)
    }
```


## 传参信息

字段后有json标识的代表需要接收的参数

用户填写表单时，前端需进行一定的验证，例如手机号为11位，名字不能过长（限定100个字）身份证号有X，不只有数字。生日和身份证号是否对的上。后端则会对重复性和邮箱有效性进行一定的验证

信息分为用户信息、账户信息以及岗位信息

前端以表单形式返回结构体字段带有json字样的值，未填写项可设为 ""，不设为nil

用户的必填项为名字、手机号、邮箱、性别，未填写项可设为 ""，不设为nil

数据库表：用户个人信息表、岗位表、用户账户表、管理员账户表

使用ID关联三个表之间的关系，前后端数据交互时要根据ID查出信息之后转移数据

### 用户表
```go
    type User struct {
        ID          string       `json:"id" gorm:"primaryKey;column:id" `
        Name        string       `json:"name" gorm:"column:name"`
        CreatedAt   time.Time    `gorm:"column:created_at"`
        State       int          `json:"state" validate:"oneof=0 1" gorm:"default:0"`
        Gender      string       `json:"gender" gorm:"column:gender"`
        Birth       string       `json:"birth"`
        Email       string       `json:"email" gorm:"column:email"`
        PhoneNumber string       `json:"phoneNumber" gorm:"column:phone"`
        ApplyID     int          `json:"applyID"`
        Account     auth.Account `json:"account"`
    }
```



### 账户信息

用户通过id（学号）+密码的方式登录

```go
    type Account struct {
        ID       string `gorm:"primarykey" json:"id"`
        Password string `json:"password"`
    }
```

### 职位表

```go
    type Job struct {
        ID        string      `json:"id" gorm:"id"`
        Name      string      `json:"name" gorm:"name"`
        ApplyList []user.User `json:"apply_list" gorm:"foreignKey:ApplyID"`
    }
```

## 后期拓展

1. 引进ai对话，帮助更好的选择自己喜欢的岗位
2. 表单填写模式不变的基础上可增加pdf提交简历
3. 岗位界面增加一级菜单二级菜单，分类更细
4. 管理员端完善，例如显示申请者之前的面评、面试进度状态
5. 添加自动审批模式，例如卡学历、工作时长，减轻管理员负担
