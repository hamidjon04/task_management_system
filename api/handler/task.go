package handler

import (
	"fmt"
	"net/http"
	"task/api/token"
	"task/model"

	"github.com/gin-gonic/gin"
)

func(h *hanlerimpl) CreateTask(c *gin.Context){
	access := c.GetHeader("Authorization")
	claim, err := token.ExtractClaimToken(access)
	if err != nil{
		h.Logger.Error(fmt.Sprintf("Claim ma'lumotlarini o'qishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Error",
		})
		return 
	}

	read := model.CreateTaskRead{}
	err = c.ShouldBindJSON(&read)
	if err != nil{
		h.Logger.Error(fmt.Sprintf("Bodydan ma'lumotlarni o'qishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Noto'g'ri ma'lumot kiritildi",
		})
		return
	}

	resp, err := h.Storage.UserService().CreateTask(&model.CreateTaskReq{
		UserId: claim.Id,
		Title: read.Title,
	})
	if err != nil{
		h.Logger.Error(fmt.Sprintf("CreateTask request error: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Error",
		})
	}
	c.JSON(http.StatusOK, resp)
}