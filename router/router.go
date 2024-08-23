package router

import (
	"classmate_reunion/controller/llm"
	"github.com/gin-gonic/gin"
)

func register(e *gin.Engine) {
	activity := llm.Activity{}
	//
	//e.Group("/openapi/v1/activity/discount/create")
	//{
	//	e.GET("/", activity.Create)
	//}

	e.GET("/openapi/v1/activity/discount/create", activity.Create)

}

func Run() {

	e := gin.Default()
	register(e)

	err := e.Run()
	if err != nil {
		return
	}

}
