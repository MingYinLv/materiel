package handle

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"materiel/src/api/v1/Log"
	"materiel/src/db/Schema"
	"materiel/src/util"
	"net/http"
)

func GetLogById(c *gin.Context) {
	if id := c.Param("id"); govalidator.IsInt(id) {
		intId, _ := govalidator.ToInt(id)
		log := Log.FindById(intId)
		if log.Id == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "该日志不存在",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "查询成功",
				"data": log,
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "id is invalid",
			"detail": []map[string]string{{"message": "must be int", "title": "id"}},
		})
	}
}

func GetLogList(c *gin.Context) {
	searchFilter := util.GetSearchFilter(c)
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "查询失败",
			})
		}
	}()
	data := Log.FindList(searchFilter)
	if len(data) == 0 {
		data = []Schema.Log{}
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "查询成功",
		"data": data,
	})
}
