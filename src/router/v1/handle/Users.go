package handle

import (
	"github.com/gin-gonic/gin"
	"materiel/src/api/v1/Users"
	"materiel/src/util"
	//"materiel/src/db"
	//"materiel/src/db/Schema"
	"net/http"
	//"strconv"
	"materiel/src/db/Schema"
	"strconv"
	"strings"
)

type TokenData struct {
	user  Schema.User
	token string
}

var tokens = make(map[int64]TokenData)

func CheckLogin(c *gin.Context) {
	token := c.GetHeader("Authorization")
	id, _ := strconv.ParseInt(util.RsaDecode(token), 10, 64)
	if data, ok := tokens[id]; ok && data.token == token {
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "请先登录",
		})
	}
}

func Info(c *gin.Context) {
	token := c.GetHeader("Authorization")
	id, _ := strconv.ParseInt(util.RsaDecode(token), 10, 64)
	c.JSON(http.StatusOK, tokens[id].user)
}

func UserLogin(c *gin.Context) {
	if val, b := c.GetPostForm("username"); b && strings.TrimSpace(val) != "" {
		u := Users.FindUserByUsername(val)
		pwd := util.GetSha256Password(c.PostForm("password"), u.Salt)
		if pwd == u.Password {
			token := util.RsaEncode(u.User_id)
			tokens[u.User_id] = TokenData{
				user: u,
				token: token,
			}
			c.JSON(http.StatusOK, gin.H{
				"msg":          "登录成功",
				"access_token": token,
				"token_type": "login",
				"expires_in": 3600,
				"refresh_token": "",

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
