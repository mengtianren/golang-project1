package model

// 字典相关
type Dict struct {
	ID     int    `json:"id"`
	Keyy   string `json:"keyy"`
	Valuee string `json:"valuee"`
	Type   string `json:"type"`
}

func (Dict) TableName() string {
	return "t_dict"
}
