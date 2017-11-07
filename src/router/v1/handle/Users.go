package handle

import (
	"github.com/gin-gonic/gin"
	"materiel/src/api/v1/Users"
	"materiel/src/util"
	//"materiel/src/db"
	//"materiel/src/db/Schema"
	"net/http"
	//"strconv"
	"strings"
)

func UserLogin(c *gin.Context) {
	if val, b := c.GetPostForm("username"); b && strings.TrimSpace(val) != "" {
		u := Users.FindUserByUsername(val)
		pwd := util.GetSha256Password(c.PostForm("password"), u.Salt)
		if pwd == u.Password {
			c.JSON(http.StatusOK, gin.H{
				"message": "登录成功",
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":  "密码错误",
				"detail": map[string]string{"message": "error", "title": "password"},
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "请输入用户名",
			"detail": map[string]string{"message": "required", "title": "username"},
		})
	}
}
