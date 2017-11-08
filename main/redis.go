package main

import (
	"github.com/gin-gonic/gin/json"
	"materiel/src/db/Schema"
	"materiel/src/redisDB"
	"strconv"
	"time"
)

type TokenData struct {
	User  Schema.User
	Token string
}

func (td *TokenData) MarshalBinary() (data []byte, err error) {
	data, err = json.Marshal(td)
	return
}

func main() {
	u := TokenData{User: Schema.User{User_id: 1, Username: "lvmingyin", Password: "79e40de7264a3ef78ddd82ede778eaeb9901abda7d457e2f8b5f1252f7cc96c4", Nickname: "吕铭印", Salt: "tdk2vwkt"}, Token: "3a811474a58ccb23c76bb4fc732cdda1224c3ef6fbc69343092affd2ca47ca21"}
	redisDB.Set(strconv.FormatInt(u.User.User_id, 10), &u, time.Hour)
}
