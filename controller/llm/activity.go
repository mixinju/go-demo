package llm

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Activity struct{}

func (a *Activity) Create(c *gin.Context) {
	poiId := c.Query("poiId")
	skuId := c.Query("skuId")
	discount := c.Query("discount")

	fmt.Println(skuId, poiId, discount)
}
