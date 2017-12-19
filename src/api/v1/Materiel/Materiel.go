package Materiel

import (
	"materiel/src/db"
	"materiel/src/db/Schema"
	//"materiel/src/util"
	"bytes"
	"database/sql"
	"fmt"
	"materiel/src/util"
	"strconv"
	"strings"
	"time"
)

//func GetMaterielList(filter util.SearchFilter) (result []Schema.Materiel) {
//	var result []Schema.Materiel
//	stms, err := db.DB.Prepare("SELECT * FROM materiel")
//	return
//}

func FindList(filter util.SearchFilter) []Schema.Materiel {
	var result []Schema.Materiel
	var buffer bytes.Buffer
	buffer.WriteString("SELECT * FROM materiel where 1=1")
	if strings.TrimSpace(filter.Keyword) != "" {
		buffer.WriteString(" and name like '%")
		buffer.WriteString(filter.Keyword)
		buffer.WriteString("%'")
	}
	if filter.SortBy == "id" || filter.SortBy == "number" {
		buffer.WriteString(" order by ")
		buffer.WriteString(filter.SortBy)
		buffer.WriteString(" ")
	}

	if filter.Order == "desc" || filter.Order == "asc" {
		buffer.WriteString(filter.Order)
	}
	if filter.Limit {
		buffer.WriteString(" limit ")
		buffer.WriteString(strconv.FormatInt((filter.Page-1)*filter.Size, 10))
		buffer.WriteString(",")
		buffer.WriteString(strconv.FormatInt(filter.Size, 10))
	}
	fmt.Println(buffer.String())
	stms, err := db.DB.Prepare(buffer.String())
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer stms.Close()
	rows, err := stms.Query()
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var materiel_id, number, create_at int64
		var name, description, change_log string
		err = rows.Scan(&materiel_id, &name, &number, &change_log, &description, &create_at)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, Schema.Materiel{materiel_id, name, number, change_log, description, create_at})
	}
	return result
}

func FindById(id int64) Schema.Materiel {
	stms, err := db.DB.Prepare("SELECT id,name,number,change_log,description,create_at FROM materiel WHERE id = ?")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	row := stms.QueryRow(id)

	var materiel_id, number, create_at int64
	var name, description, change_log string

	err = row.Scan(&materiel_id, &name, &number, &change_log, &description, &create_at)
	if err == sql.ErrNoRows {
		stms.Close()
		return Schema.Materiel{Id: 0}
	} else if err != nil {
		panic(err.Error())
	}
	stms.Close()
	return Schema.Materiel{materiel_id, name, number, change_log, description, create_at}
}

func AddMateriel(materiel *Schema.Materiel, log *Schema.Log) int64 {
	tx, _ := db.DB.Begin()
	stms, err := tx.Prepare("insert into materiel(name, number, change_log, description, create_at) values(?, ?, ?, ?, ?)")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	unix := time.Now().Unix()
	result, err := stms.Exec(materiel.Name, materiel.Number, materiel.Number, materiel.Description, unix)
	if err != nil {
		panic(err.Error())
	}

	insertId, err := result.LastInsertId()

	if err != nil {
		tx.Rollback()
		panic(err.Error())
	}
	stms.Close()

	stms, err = tx.Prepare("insert into log(materiel_id, number,quantity, type,operator,operate_time, remark, create_at) values(?, ?,?,?,?, ?, ?, ?)")
	if err != nil {
		tx.Rollback()
		panic(err.Error())
	}

	result, err = stms.Exec(insertId, log.Number, log.Quantity, log.Type, log.Operator, log.OperateTime, log.Remark, unix)
	if err != nil {
		tx.Rollback()
		panic(err.Error())
	}
	insertLogId, err := result.LastInsertId()
	log.Id = insertLogId
	log.CreateAt = unix
	log.MaterielId = insertId

	if err != nil {
		tx.Rollback()
		panic(err.Error())
	}

	stms.Close()
	tx.Commit()
	materiel.Id = insertId
	materiel.CreateAt = unix
	return insertId
}

/*
*	更新物料数量，记录日志
 */
func UpdateMateriel(materiel *Schema.Materiel, log *Schema.Log) int64 {
	// 开启事务
	tx, _ := db.DB.Begin()
	// 修改物料数量
	stms, err := tx.Prepare("update materiel set name=?,number=?,change_log=CONCAT(change_log, \",\" ,?),description=? where id=?")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	result, err := stms.Exec(materiel.Name, materiel.Number, materiel.Number, materiel.Description, materiel.Id)
	if err != nil {
		tx.Rollback()
		panic(err.Error())
	}
	stms.Close()

	affectedRow, err := result.RowsAffected()

	if err != nil {
		tx.Rollback()
		panic(err.Error())
	}

	// 保存日志
	stms, err = tx.Prepare("insert into log(materiel_id, number, type, operator, operate_time, remark, create_at) values(?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		tx.Rollback()
		panic(err.Error())
	}

	t := time.Now()
	result, err = stms.Exec(log.MaterielId, log.Number, log.Type, log.Operator, log.OperateTime, log.Remark, t.Unix())

	if err != nil {
		tx.Rollback()
		panic(err.Error())
	}
	insertLogId, err := result.LastInsertId()
	log.Id = insertLogId
	log.CreateAt = t.Unix()
	if err != nil {
		tx.Rollback()
		panic(err.Error())
	}

	stms.Close()
	tx.Commit()

	return affectedRow
}

func DeleteMaterielById(id int64) (int64, error) {
	tx, _ := db.DB.Begin()
	stms, err := tx.Prepare("delete from materiel WHERE id = ?")
	if err != nil {
		tx.Rollback()
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer stms.Close()

	result, err := stms.Exec(id)
	if err != nil {
		tx.Rollback()
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	stms, err = tx.Prepare("delete from logs WHERE materiel_id = ?")
	if err != nil {
		tx.Rollback()
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	_, err = stms.Exec(id)
	if err != nil {
		tx.Rollback()
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	tx.Commit()
	return result.RowsAffected()
}
