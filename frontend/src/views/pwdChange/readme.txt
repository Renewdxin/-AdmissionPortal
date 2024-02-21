1.class
pwdChange
    main
        demo-tabs
            el-tabs__content(发送短信验证码的途径)
                before
                    elform(el-form)
                after
                    elform
        demo-tabs
            el-tabs__content(旧密码验证的途径)
                before
                    elform

2.部分函数
发送验证码：sendCode
验证验证码：judgeCode
用验证码方式修改密码：changePwdCode
用旧密码方式修改密码：changePwdOld

