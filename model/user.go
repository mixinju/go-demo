package model

type User struct {
	Id         uint        // userId
	Name       string      // 姓名
	OpenId     string      // 微信openid
	Phone      uint64      // 电话号码
	Educations []Education // 教育经历 Education 组成的数组字符串
	City       string      // 城市
	Sex        bool        // 性别 true:男 false:女
	LoveStatus string      // 恋爱状态
}
