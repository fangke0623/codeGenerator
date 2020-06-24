package discuss

import "time"

type discuss struct {

//聊天室id
    discuss_id     int64 
    //聊天室名称
    discuss_title    string 
    //发起人id
    user_id    string 
    //加入方式 1.自由加入 2.邀请加入
    visible_type     int64 
    //话题状态 1-正常 2-删除
    status     int64 
    //创建时间 
    create_time     time.Time 
    //修改时间
    modify_time     time.Time
    //修改人
    modify_id    string 
    
}