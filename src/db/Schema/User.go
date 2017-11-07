package Schema

type User struct {
	User_id  int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Salt     string `json:"salt"`
}
