package data

import (
	"github.com/gin-gonic/gin"
	. "go-crawler/crawler/backend/handler"
)

// /admin/data/?type=?&id=?
func DeleteData(c *gin.Context) {

	did := c.DefaultQuery("id", "0")

	typed := c.DefaultQuery("type", "0")


	item := Item{
		Type: typed,
		Id: did,
	}

	err := Delete(index, item)

	SendResponse(c, err, nil)
}
