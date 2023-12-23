package model

import "time"

type Post struct {
	Id        uint64
	Content   string    // 文章内容
	CreatedAt time.Time // 创建(发布)时间
	UserId    uint64 
	Like      uint32
}
