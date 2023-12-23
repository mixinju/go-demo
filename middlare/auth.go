package middlare

import "github.com/gin-gonic/gin"

func Auth(c *gin.Context) {
	// 是否注册. 未注册,重定向到注册页面
	openId := c.GetHeader("OpenID")
	if len(openId) == 0 {
		// 登陆失败
		c.Abort()
	}

	// 注册--> 通过 OpenID 置换userId
	// 前后端交互使用userId 标识用户身份

}
