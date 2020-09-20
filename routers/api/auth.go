package api

import (
	"github.com/NonePro/blog/models"
	"github.com/NonePro/blog/pkg/err"
	"github.com/NonePro/blog/pkg/logging"
	"github.com/NonePro/blog/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := err.INVALID_PARAMS
	if ok {
		isExist := models.CheckAuth(username, password)
		if isExist {
			token, e := util.GenerateToken(username, password)
			if e != nil {
				code = err.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token

				code = err.SUCCESS
			}

		} else {
			code = err.ERROR_AUTH
		}
	} else {
		for _, e := range valid.Errors {
			logging.Info(e.Key, e.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  err.GetMsg(code),
		"data": data,
	})
}
