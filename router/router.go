package router

import (
	"fmt"
	"gin-api/model"
	"log"
	"net/http"
	"strconv"

	dao "gin-api/dao"
	"github.com/gin-gonic/gin"
)

func StartRouter() {
	router := gin.Default()

	router.GET("/message/:id", func(c *gin.Context) {
		id := c.Param("id")
		u64, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			fmt.Println(err)
		}
		n := uint(u64)
		result := dao.GetMessage(n)
		c.IndentedJSON(http.StatusOK, result)
	})

	router.GET("/message", func(c *gin.Context) {
		result := dao.GetAllMessage()
		c.IndentedJSON(http.StatusOK, result)
	})

	router.POST("/message", func(c *gin.Context) {
		var creatreMessage model.Message

		if err := c.BindJSON(&creatreMessage); err != nil {
			log.Fatal("BIND JSON ERROR")
			return
		}
		result := dao.CreateMessage(creatreMessage.MSG)
		//result := db.SaveMessage(newMessage)
		c.IndentedJSON(http.StatusCreated, result)
	})

	router.PUT("/message", func(c *gin.Context) {
		var updateMessage model.Message

		if err := c.BindJSON(&updateMessage); err != nil {
			log.Fatal("BIND JSON ERROR")
			return
		}

		result := dao.UpdateMessage(updateMessage)
		c.IndentedJSON(http.StatusAccepted, result)
	})

	router.DELETE("/message", func(c *gin.Context) {
		var deleteMessage model.Message

		if err := c.BindJSON(&deleteMessage); err != nil {
			log.Fatal("BIND JSON ERROR")
			return
		}

		result := dao.DeleteMessage(deleteMessage.ID)
		c.IndentedJSON(http.StatusAccepted, result)
	})

	router.Run("localhost:8080")
}
