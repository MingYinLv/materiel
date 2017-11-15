package Schema

type Materiel struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Number      int64  `json:"number"`
	ChangeLog	string `json:"change_log"`
	Description string `json:"description"`
	CreateAt    int64  `json:"create_at"`
}
