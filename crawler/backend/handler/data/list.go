package data

import (
	"github.com/gin-gonic/gin"
	"strconv"
	. "go-crawler/crawler/backend/handler"
)

func DataList(c *gin.Context) {
	q := c.DefaultQuery("q", "")
	from, err := strconv.Atoi(c.DefaultQuery("from", "0"))
	if err != nil{
		SendResponse(c, err, nil)
	}

	result,err := GetSearchResult(q, from)

	SendResponse(c, err, result)
}