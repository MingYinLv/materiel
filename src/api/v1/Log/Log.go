package Log

import (
	"materiel/src/db"
	"materiel/src/db/Schema"
)

func AddLog(log Schema.Log) (int64, error) {
	stms, err := db.DB.Prepare("insert into log(id,materiel_id, number, type, remark) values(?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}

	result, err := stms.Exec(11,log.Materiel_id, log.Number, log.Type, log.Remark)
	if err != nil {
		return 0, err
	}
	stms.Close()
	return result.LastInsertId()
}
