package llm

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Activity struct{}

func (a *Activity) Create(c *gin.Context) {
	poiId := c.Query("poiId")
	skuId := c.Query("skuId")
	discount := c.Query("discount")

	fmt.Println(skuId, poiId, discount)

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
	})
}
