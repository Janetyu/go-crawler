package data

import (
	"github.com/gin-gonic/gin"
	. "go-crawler/crawler/backend/handler"
	"go-crawler/crawler/backend/pkg/errno"
)



func SaveOrUpdateData(c *gin.Context) {
	var r Item
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	err := SaveOrUpdate(index, r)

	// Show the user information.
	SendResponse(c, err, nil)
}
