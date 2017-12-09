package Log

import (
	"bytes"
	"database/sql"
	"fmt"
	"materiel/src/db"
	"materiel/src/db/Schema"
	"materiel/src/util"
	"strconv"
	"strings"
)

func FindById(id int64) Schema.Log {
	stms, err := db.DB.Prepare("SELECT * FROM log WHERE id = ?")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	row := stms.QueryRow(id)

	var log_id, materiel_id, number, quantity, log_type, create_at, operate_time int64
	var operator, remark string

	err = row.Scan(&log_id, &materiel_id, &number, &quantity, &log_type, &operator, &operate_time, &remark, &create_at)
	if err == sql.ErrNoRows {
		stms.Close()
		return Schema.Log{Id: 0}
	} else if err != nil {
		panic(err.Error())
	}
	stms.Close()
	return Schema.Log{log_id, materiel_id, number, quantity, log_type, operator, operate_time, remark, create_at}
}

func FindList(filter util.SearchFilter) []Schema.Log {
	var result []Schema.Log
	var buffer bytes.Buffer
	buffer.WriteString("SELECT * FROM log where 1=1")
	if strings.TrimSpace(filter.Keyword) != "" {
		buffer.WriteString(" and operator like '%")
		buffer.WriteString(filter.Keyword)
		buffer.WriteString("%'")
	}
	if filter.Id > 0 {
		buffer.WriteString(" and materiel_id=")
		buffer.WriteString(strconv.FormatInt(filter.Id, 10))
	}
	if filter.Type > Schema.ALL {
		buffer.WriteString(" and type=")
		buffer.WriteString(strconv.FormatInt(filter.Type, 10))
	}
	if filter.SortBy == "id" || filter.SortBy == "number" {
		buffer.WriteString(" order by ")
		buffer.WriteString(filter.SortBy)
	}

	if filter.Order == "desc" || filter.Order == "asc" {
		buffer.WriteString(" ")
		buffer.WriteString(filter.Order)
	}
	if filter.Limit{
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
		var log_id, materiel_id, number, quantity, log_type, create_at, operate_time int64
		var operator, remark string

		err = rows.Scan(&log_id, &materiel_id, &number, &quantity, &log_type, &operator, &operate_time, &remark, &create_at)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, Schema.Log{log_id, materiel_id, number, quantity, log_type, operator, operate_time, remark, create_at})
	}
	return result
}
