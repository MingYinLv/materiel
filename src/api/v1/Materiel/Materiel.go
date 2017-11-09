package Materiel

import (
	"fmt"
	"materiel/src/db"
	"materiel/src/db/Schema"
)

func FindById(id int64) Schema.Materiel {
	stms, err := db.DB.Prepare("SELECT id,name,number,description FROM users WHERE id = ?")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	row := stms.QueryRow(id)

	var materiel_id, number int64
	var name, description string

	err = row.Scan(&materiel_id, &name, &number, &description)
	stms.Close()
	return Schema.Materiel{materiel_id, name, number, description}
}

func AddMateriel(materiel Schema.Materiel) (int64, error) {
	tx, _ := db.DB.Begin()
	stms, err := tx.Prepare("insert into materiel(name, number, description) values(?, ?, ?)")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	result, err := stms.Exec(materiel.Name, materiel.Number, materiel.Description)
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

	stms, err = tx.Prepare("insert into log(materiel_id, number, type, remark) values(?, ?, ?, ?)")
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

	result, err = stms.Exec(log.Materiel_id, log.Number, log.Type, log.Remark)
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
