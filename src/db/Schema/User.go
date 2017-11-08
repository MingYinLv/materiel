package Schema

import (
	"encoding/json"
)

type User struct {
	User_id  int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"_"`
	Nickname string `json:"nickname"`
	Salt     string `json:"_"`
}

func (u *User) MarshalBinary() (data []byte, err error) {
	data, err = json.Marshal(u)
	return
}
