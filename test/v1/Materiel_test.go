package v1

import (
	"github.com/astaxie/beego/logs"
	"materiel/src/api/v1/Materiel"
	"materiel/src/db/Schema"
	"testing"
)

func Test_AddMateriel(t *testing.T) {
	insertId, err := Materiel.AddMateriel(Schema.Materiel{
		Name:   "笔记本",
		Number: 100,
		Description: "云定制的笔记本",
	})

	if err != nil {
		logs.Error("添加物料失败")
	} else {
		logs.Error("添加物料成功", insertId)
	}
}
