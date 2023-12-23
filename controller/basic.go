package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Basic 提供基础能力 登陆和注册, 图片上传
type Basic struct{}

func (b *Basic) Login(c *gin.Context) {
	c.JSON(http.StatusOK,
		gin.H{
			"code":    0,
			"message": "ok",
		})
}

func (b *Basic) Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "ok"})
}

func (b *Basic) Upload(c *gin.Context) {

}

func (b *Basic) Banner(c *gin.Context) {

}
