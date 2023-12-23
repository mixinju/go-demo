package service

type User struct{}

func (u *User) QueryIdByOpenId(openId string) uint {
	return 1
}
