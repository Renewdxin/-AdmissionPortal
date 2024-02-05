## 地址
git仓库地址：先提交到各自的分支

后端：back_branch

前端：front_branch


路由
```go
func NewRouter() *gin.Engine {
    r := gin.New()
    // home page 主要为各个入口，打算直接一个html定死
    r.POST("/home")

    // account setting
    apiAccount := r.Group("/auth")
    apiAccount.Use(middleware.JWTHandler())
    {
       // 密码+手机号
       apiAccount.POST("/login")
       // 主要功能为更改cokkie中的max-age，退出登录状态
       apiAccount.POST("/logout")
       // 用户注册
       // 1. 是否已被注册
       // 2. 验证码
       // 3. 先用手机号+密码注册
       // 后续用户进入个人信息页自行完善信息
       apiAccount.POST("/signup")
       // 忘记密码： 需要进行手机号或者邮箱验证码验证过后，方能修改密码；这里应该还可以分出链接，通过邮箱还是手机号验证码
       apiAccount.POST("/password/forget")
       // 旧密码核对正确即可
       apiAccount.POST("/password/change")
    }

    // personal info
    apiProfile := r.Group("/profile")
    apiProfile.Use(middleware.JWTHandler())
    {
       // GetUserInfo 通过id得到用户信息，返回项为姓名、性别、出生日期、邮箱、手机号
        apiProfile.GET("/Info/:id", handler.GetUserInfo)
        // DeleteUser 用户在删除前需要进行手机验证码验证才能删除
        apiProfile.DELETE("/delete/:id", handler.DeleteUser)
        // 更新用户信息，仅限手机号、邮箱
        apiProfile.PUT("/update/:id", handler.UpdateUserInfo)
        // 查询是否被录取，这里前端通过后端传回的状态码进行输出，录取还是未录取
        apiProfile.GET("/status/:id", handler.GetUserStatus)
    }
    
    // admin（账号定死）
    apiAdmin := r.Group("/admin")
    apiAdmin.Use() // 使用JWT中间件进行管理员身份验证
    {
        // 管理员仪表板或主页，后端解析一个html文件
        apiAdmin.GET("/dashboard")
    
        // 查看所有职位发布（管理员）
        apiAdmin.GET("/jobs")
    
        // 查看职位详情（管理员）
        apiAdmin.GET("/job/:jobID")
    
        // 查看职位申请状态（管理员）
        apiAdmin.GET("/applications/:jobID")
    
        // 审批或拒绝职位申请（管理员）
        apiAdmin.PUT("/application/:appID")
    }
    
    // recruitment
    // 职位就4个 前后端、两个客户端
    apiJob := r.Group("/recruitment")
    apiJob.Use()
    {
        // 岗位总览
        apiJob.GET("/jobs")
        // 岗位详细信息，例如应聘要求
        apiJob.GET("/job/:id")
        // 申请投递，需确保登录状态，若未登录则先跳转至登录页面；
        // 后期拓展为可pdf提交简历（原填写框不变）
        apiJob.POST("/job/:id/apply")
    }
    return r
}
```


## 传参信息
用户填写表单时，前端需进行一定的验证，例如手机号为11位，名字不能过长（限定100个字）身份证号有X，不只有数字。生日和身份证号是否对的上。后端则会对重复性和邮箱有效性进行一定的验证

信息分为用户信息、账户信息以及岗位信息

前端以表单形式返回结构体字段带有json字样的值，未填写项可设为 ""，不设为nil

用户的必填项为名字、手机号、邮箱、性别，未填写项可设为 ""，不设为nil

```go
type User struct {
    ID          uint   `gorm:"primarykey" json:"id"`
    Name        string `json:"name"`
    CreatedAt   time.Time
    State       int    `json:"state" gorm:"type:tinyint" validate:"oneof=0 1"`
    Gender      string `json:"gender"`
    Birth       string `json:"birth"`
    Email       string `json:"id"`
    PhoneNumber string `json:"phoneNumber"`
    Account     account.Account
}   
```

字段后有json标识的代表需要接收的参数

id可前端设定一并传回，也可以后端自动生成，id为用户标识符，不可重复

账户信息

```go
type Account struct {
    ID       string `gorm:"primarykey" json:"id"`
    Password string `json:"password"`
}
```

账户信息通过学号登录

id前端或者后端自动生成

```go
type Job struct {
    ID string `json:"id" gorm:"id"`
    Name string `json:"name" gorm:"name"`
}
```

数据库表：用户个人信息表，岗位表，账户密码表

使用ID关联三个表之间的关系，前后端数据交互时要根据ID查出信息之后转移数据

后期拓展
1. 引进ai对话，帮助更好的选择自己喜欢的岗位
2. 表单填写模式不变的基础上可增加pdf提交简历
3. 岗位界面增加一级菜单二级菜单，分类更细
4. 管理员端完善，例如显示申请者之前的面评、面试进度状态
5. 添加自动审批模式，例如卡学历、工作时长，减轻管理员负担
