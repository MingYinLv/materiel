package Schema

type Materiel struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Number      int64  `json:"number"`
	Description string `json:"description"`
}
