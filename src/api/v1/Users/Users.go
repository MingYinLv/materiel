package Users

import "materiel/src/db/Schema"
import (
	"materiel/src/db"
	"time"
)

func FindUserByUsername(name string) Schema.User {
	stms, err := db.DB.Prepare("SELECT id,username,password,nickname,salt,create_at FROM users WHERE username = ?")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	row := stms.QueryRow(name)

	var user_id, create_at int64
	var username, password, nickname, salt string

	err = row.Scan(&user_id, &username, &password, &nickname, &salt, &create_at)
	if err != nil {
		panic(err.Error())
	}
	stms.Close()
	return Schema.User{user_id, username, password, nickname, salt, create_at}
}

func FindUserById(id int64) Schema.User {
	stms, err := db.DB.Prepare("SELECT id,username,password,nickname,salt,create_at FROM users WHERE id = ?")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	row := stms.QueryRow(id)

	var user_id, create_at int64
	var username, password, nickname, salt string

	err = row.Scan(&user_id, &username, &password, &nickname, &salt, &create_at)
	if err != nil {
		panic(err.Error())
	}
	stms.Close()
	return Schema.User{user_id, username, password, nickname, salt, create_at}
}

func AddUser(user Schema.User) (int64, error) {
	stms, err := db.DB.Prepare("insert into users(username, nickname, password, salt, create_at) values(?, ?, ?, ?, ?)")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	t := time.Now()
	result, err := stms.Exec(user.Username, user.Nickname, user.Password, user.Salt, t.Unix())
	stms.Close()
	return result.LastInsertId()
}

func UpdateUser(user Schema.User) (int64, error) {
	stms, err := db.DB.Prepare("update users set username=?,password=?,nickname=?,salt=? where id=?")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	result, err := stms.Exec(user.Username, user.Password, user.Nickname, user.Salt, user.User_id)
	stms.Close()
	return result.RowsAffected()
}

func DeleteUserById(id int64) (int64, error) {
	stms, err := db.DB.Prepare("delete from users WHERE id = ?")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	result, err := stms.Exec(id)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	stms.Close()
	return result.RowsAffected()
}
