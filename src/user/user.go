package user

import "time"

type User struct {

//
    id    string 
    //邮箱
    email    string 
    //用户名
    user_name    string 
    //密码
    password    string 
    //创建时间
    create_time     time.Time 
    //修改时间
    modify_time     time.Time 
    //昵称
    nickname    string 
    //手机号
    mobile    string 
    
}