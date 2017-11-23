package handle

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"materiel/src/api/v1/Materiel"
	"materiel/src/db/Schema"
	"materiel/src/util"
	"net/http"
	"time"
)

func UpdateMateriel(c *gin.Context) {
	if id := c.Param("id"); govalidator.IsInt(id) {
		if util.RequiredValidate(c, []string{
			"name", "number", "operator", "operate_time",
		}) && util.IntValidate(c, []string{
			"number", "type",
		}) {
			materielId, _ := govalidator.ToInt(id)
			find := Materiel.FindById(materielId)
			if find.Id == 0 {
				c.JSON(http.StatusNotFound, gin.H{
					"error": "该物料不存在",
				})
				return
			}

			operate_time, err := time.Parse("2006-01-02", c.PostForm("operate_time"))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error":  "日期格式错误",
					"detail": []map[string]string{{"message": "format error", "title": "operate_time"}},
				})
				return
			}
			name := c.PostForm("name")
			change_log := c.PostForm("change_log")
			description := c.PostForm("description")
			number, _ := govalidator.ToInt(c.PostForm("number"))
			operator := c.PostForm("operator")
			remark := c.PostForm("remark")
			logType, _ := govalidator.ToInt(c.PostForm("type"))
			quantity := number

			if number < 0 {
				c.JSON(http.StatusBadRequest, gin.H{
					"error":  "修改数量不能小于0",
					"detail": []map[string]string{{"message": "number out of bounds", "title": "number"}},
				})
				return
			}

			if logType == Schema.OUT && find.Number-number < 0 {
				c.JSON(http.StatusBadRequest, gin.H{
					"error":  "库存不足，无法出库",
					"detail": []map[string]string{{"message": "number out of bounds", "title": "number"}},
				})
				return
			}

			if logType == Schema.OUT {
				quantity = find.Number - number
			} else if logType == Schema.IN {
				quantity = find.Number + number
			}

			materiel := Schema.Materiel{
				Id:          materielId,
				Name:        name,
				Number:      quantity,
				Description: description,
				CreateAt:    find.CreateAt,
				ChangeLog:   fmt.Sprintf("%s,%d", change_log, quantity),
			}
			log := Schema.Log{
				MaterielId:  materielId,
				Number:      number,
				Type:        logType,
				OperateTime: operate_time.Unix(),
				Operator:    operator,
				Remark:      remark,
			}
			defer func() {
				if err := recover(); err != nil {
					fmt.Println(err)
					c.JSON(http.StatusBadRequest, gin.H{
						"error": "物料修改失败",
					})
				}
			}()
			Materiel.UpdateMateriel(&materiel, &log)
			c.JSON(http.StatusOK, gin.H{
				"msg": "操作成功",
				"data": gin.H{
					"materiel": materiel,
					"log":      log,
				},
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "id is invalid",
			"detail": []map[string]string{{"message": "must be int", "title": "id"}},
		})
	}
}

func GetMaterielById(c *gin.Context) {
	if id := c.Param("id"); govalidator.IsInt(id) {
		intId, _ := govalidator.ToInt(id)
		materiel := Materiel.FindById(intId)
		if materiel.Id == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "该物料不存在",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "查询成功",
				"data": materiel,
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "id is invalid",
			"detail": []map[string]string{{"message": "must be int", "title": "id"}},
		})
	}
}

func GetMaterielList(c *gin.Context) {
	searchFilter := util.GetSearchFilter(c)
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "查询失败",
			})
		}
	}()

	data := Materiel.FindList(searchFilter)
	if len(data) == 0 {
		data = []Schema.Materiel{}
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "查询成功",
		"data": data,
	})
}

func AddMateriel(c *gin.Context) {
	if util.RequiredValidate(c, []string{
		"name", "number", "operator", "operate_time",
	}) {
		name := c.PostForm("name")
		number := c.PostForm("number")
		operator := c.PostForm("operator")
		operate_time, err := time.Parse("2006-01-02", c.PostForm("operate_time"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":  "日期格式错误",
				"detail": []map[string]string{{"message": "format error", "title": "operate_time"}},
			})
			return
		}
		if !govalidator.IsInt(number) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":  "物料数量只能为整数",
				"detail": []map[string]string{{"message": "must int", "title": "number"}},
			})
			return
		}

		numberInt, _ := govalidator.ToInt(number)
		log := Schema.Log{
			Number:      numberInt,
			Type:        Schema.INSERT,
			OperateTime: operate_time.Unix(),
			Operator:    operator,
			Remark:      c.PostForm("remark"),
		}
		materiel := Schema.Materiel{Name: name, Number: numberInt, Description: c.PostForm("description")}
		defer func() {
			if err := recover(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "物料添加失败",
				})
			}
		}()
		Materiel.AddMateriel(&materiel, &log)
		c.JSON(http.StatusOK, gin.H{
			"msg": "查询成功",
			"data": gin.H{
				"materiel": materiel,
				"log":      log,
			},
		})
	}
}
