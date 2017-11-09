package Materiel

import (
	"fmt"
	"materiel/src/db"
	"materiel/src/db/Schema"
	//"materiel/src/util"
	"time"
)

//func GetMaterielList(filter util.SearchFilter) (result []Schema.Materiel) {
//	var result []Schema.Materiel
//	stms, err := db.DB.Prepare("SELECT * FROM materiel")
//	return
//}

func FindById(id int64) Schema.Materiel {
	stms, err := db.DB.Prepare("SELECT id,name,number,description,create_at FROM users WHERE id = ?")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	row := stms.QueryRow(id)

	var materiel_id, number, create_at int64
	var name, description string

	err = row.Scan(&materiel_id, &name, &number, &description, &create_at)
	stms.Close()
	return Schema.Materiel{materiel_id, name, number, description, create_at}
}

func AddMateriel(materiel Schema.Materiel) (int64, error) {
	tx, _ := db.DB.Begin()
	stms, err := tx.Prepare("insert into materiel(name, number, description, create_at) values(?, ?, ?, ?)")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	t := time.Now()
	result, err := stms.Exec(materiel.Name, materiel.Number, materiel.Description, t.Unix())
	if err != nil {
		tx.Rollback()
		panic(err.Error())
	}

	insertId, err := result.LastInsertId()

	if err != nil {
		tx.Rollback()
		panic(err.Error())
	}
	stms.Close()

	stms, err = tx.Prepare("insert into log(materiel_id, number, type, remark, create_at) values(?, ?, ?, ?, ?)")
	if err != nil {
		tx.Rollback()
		panic(err.Error())
	}

	log := Schema.Log{
		Materiel_id: insertId,
		Number:      materiel.Number,
		Type:        Schema.INSERT,
		Remark:      fmt.Sprintf("新增物料: %s,数量:%d", materiel.Name, materiel.Number),
	}

	result, err = stms.Exec(log.Materiel_id, log.Number, log.Type, log.Remark, t.Unix())
	if err != nil {
		tx.Rollback()
		panic(err.Error())
	}
	_, err = result.LastInsertId()

	if err != nil {
		tx.Rollback()
		panic(err.Error())
	}

	stms.Close()
	tx.Commit()
	return insertId, err
}

/*
*	更新物料数量，记录日志
 */
func UpdateMateriel(materiel Schema.Materiel, log Schema.Log) (int64, error) {
	// 开启事务
	tx, _ := db.DB.Begin()
	// 修改物料数量
	stms, err := tx.Prepare("update materiel set number=? where id=?")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	result, err := stms.Exec(materiel.Number, materiel.Id)
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
	stms, err = tx.Prepare("insert into log(materiel_id, number, type, remark, create_at) values(?, ?, ?, ?, ?)")
	if err != nil {
		tx.Rollback()
		panic(err.Error())
	}

	t := time.Now()
	result, err = stms.Exec(log.Materiel_id, log.Number, log.Type, log.Remark, t.Unix())

	if err != nil {
		tx.Rollback()
		panic(err.Error())
	}
	_, err = result.LastInsertId()

	if err != nil {
		tx.Rollback()
		panic(err.Error())
	}

	stms.Close()
	tx.Commit()

	return affectedRow, err
}
