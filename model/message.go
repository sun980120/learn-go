package model

type Message struct {
	ID  uint   `json:"id" gorm:"primary_key"`
	MSG string `json:"msg"`
}
