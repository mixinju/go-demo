package model

type Image struct {
	Id     uint64
	Url    string // COS 地址
	Tag    string // 自定义识别 Tag
	FaceId string // 腾讯云人脸识别唯一ID
}
