package model

type Message struct {
	ID  uint   `json:"id" gorm:"primary_key"`
	Msg string `json:"msg"`
}
