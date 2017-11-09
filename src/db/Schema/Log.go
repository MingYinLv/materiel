package Schema

const (
	INSERT = iota
	UPDATE
	IN
	OUT
)

type Log struct {
	Id          int64  `json:"id"`
	Materiel_id int64  `json:"materiel_id"`
	Number      int64  `json:"number"`
	Type        int64  `json:"type"`
	Remark      string `json:"remark"`
}
