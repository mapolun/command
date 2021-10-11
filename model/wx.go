package model

type WxLogin struct {
	Id   int64 `gorm:"primaryKey"`
	Name string
}

func (WxLogin) TableName() string {
	return "wx_login"
}
