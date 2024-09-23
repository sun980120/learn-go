package dao

import (
	"gin-api/db"
	"gin-api/model"
)

func CreateMessage(msg string) string {
	conn := db.Connect()
	message := model.Message{Msg: msg}
	conn.Create(&message)
	return "create message"
}

func GetMessage(id uint) model.Message {
	conn := db.Connect()
	message := model.Message{}
	conn.Where("id = ?", id).Find(&message)
	return message
	//selectDynStmt := `SELECT * FROM Message where "id" = $1`
	//rows, err := db.Select()
	//message := conn.Where()
}

func GetAllMessage() []model.Message {
	conn := db.Connect()
	var message []model.Message
	conn.Find(&message)
	return message
}

func UpdateMessage(msg model.Message) string {
	conn := db.Connect()
	conn.Model(&model.Message{}).Where("id = ?", msg.ID).Update("msg", msg.Msg)
	return "update success"
}

func DeleteMessage(id uint) string {
	conn := db.Connect()
	conn.Where("id = ?", id).Delete(&model.Message{})
	return "delete success"
}
