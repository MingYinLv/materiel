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
	"strings"
	//"materiel/src/redisDB"
	"encoding/json"
	"github.com/go-redis/redis"
	"materiel/src/redisDB"
	"strconv"
	"time"
)

type TokenData struct {
	User          Schema.User `json:"user"`
	Token         string      `json:"token"`
	Refresh_Token string      `json:"refresh_token"`
}

func (td *TokenData) MarshalBinary() (data []byte, err error) {
	data, err = json.Marshal(td)
	return
}

func CheckLogin(c *gin.Context) {
	token := c.GetHeader("Authorization")
	result := redisDB.Get(util.RsaDecode(token))
	if result == redis.Nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "请先登录",
		})
		return
	}

	var td TokenData
	json.Unmarshal([]byte(result.(string)), &td)

	if td.Token == token {
		c.Set("td", td)
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "token已过期",
		})
	}
}

func Info(c *gin.Context) {
	if td, ok := c.Get("td"); ok {
		u := td.(TokenData)
		c.JSON(http.StatusOK, u.User)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "错误请求",
		})
	}
}

func RefreshToken(c *gin.Context) {
	if td, b := c.Get("td"); b {
		u := td.(TokenData)
		refresh_token, ok := c.GetPostForm("refresh_token")
		if ok && strings.TrimSpace(refresh_token) != "" {
			if refresh_token == u.Refresh_Token {
				token := util.RsaEncode(u.User.User_id)
				refresh_token := util.RsaEncode(u.User.User_id)
				redisDB.Set(strconv.FormatInt(u.User.User_id, 10), &TokenData{
					Token:         token,
					Refresh_Token: refresh_token,
					User:          u.User,
				}, time.Hour)
				c.JSON(http.StatusOK, gin.H{
					"msg":           "token获取成功",
					"access_token":  token,
					"token_type":    "login",
					"expires_in":    3600,
					"refresh_token": refresh_token,
				})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"error":  "refresh_token无效",
					"detail": []map[string]string{{"message": "invalid", "title": "refresh_token"}},
				})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":  "refresh_token不能为空",
				"detail": []map[string]string{{"message": "required", "title": "refresh_token"}},
			})
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":  "token不能为空",
			"detail": []map[string]string{{"message": "required", "title": "token"}},
		})
	}
}

func UserLogin(c *gin.Context) {
	if val, b := c.GetPostForm("username"); b && strings.TrimSpace(val) != "" {
		u := Users.FindUserByUsername(val)
		pwd := util.GetSha256Password(c.PostForm("password"), u.Salt)
		if pwd == u.Password {
			token := util.RsaEncode(u.User_id)
			refresh_token := util.RsaEncode(u.User_id)
			redisDB.Set(strconv.FormatInt(u.User_id, 10), &TokenData{
				Token:         token,
				Refresh_Token: refresh_token,
				User:          u,
			}, time.Hour)
			c.JSON(http.StatusOK, gin.H{
				"msg":           "登录成功",
				"access_token":  token,
				"token_type":    "login",
				"expires_in":    3600,
				"refresh_token": refresh_token,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":  "密码错误",
				"detail": []map[string]string{{"message": "error", "title": "password"}},
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "请输入用户名",
			"detail": []map[string]string{{"message": "required", "title": "username"}},
		})
	}
}
