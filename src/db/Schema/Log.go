package Schema

const (
	INSERT = iota
	UPDATE
	IN
	OUT
)

type Log struct {
	Id          int64  `json:"id"`
	MaterielId  int64  `json:"materiel_id"`
	Number      int64  `json:"number"`
	Type        int64  `json:"type"`
	Operator    string `json:"operator"`
	OperateTime int64  `json:"operate_time"`
	Remark      string `json:"remark"`
	CreateAt    int64  `json:"create_at"`
}
