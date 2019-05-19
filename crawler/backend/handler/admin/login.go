package admin

import (
	"github.com/gin-gonic/gin"

	. "go-crawler/crawler/backend/handler"

	"go-crawler/crawler/backend/pkg/errno"
	"go-crawler/crawler/backend/model"
)

// 账号+密码登录
func AdminLogin(c *gin.Context) {
	var r LoginRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	admin, err := model.GetAdmin(r.Username)
	if err != nil {
		SendResponse(c, errno.ErrAdminNotFound, nil)
		return
	}

	err = admin.Compare(r.Password)
	if err != nil {
		SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}

	rsp := LoginResponse{
		Id:       admin.Id,
		Username: admin.Username,
		RoleId:   admin.RoleId,
		Name:    admin.Name,
	}

	SendResponse(c, nil, rsp)
}
