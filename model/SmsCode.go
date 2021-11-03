package model

type SmsCode struct {
	Id int64 `xorm:"pk autoincr" json:"id"`
	Phone string `xorm:"varchar(11)" json:"phone"`
	Code string `xorm:"varchar(6)" json:"code"`
	CreateTime int64 `xorm:"bigint" json:"create_time"`
}
