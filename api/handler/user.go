package handler

import (
	"fmt"
	"net/http"
	"task/api/token"
	"task/model"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func(h *hanlerimpl) Register(c *gin.Context){
	req := model.RegisterReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil{
		h.Logger.Error(fmt.Sprintf("Ma'lumotlarni bodydan o'qishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Xato ma'lumotlar kiritildi",
		})
		return
	}
	
	hashpassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil{
		h.Logger.Error(fmt.Sprintf("Passwordni hashlashda xatolik: %v", err))
	}
	req.Password = string(hashpassword)

	resp, err := h.Storage.UserService().Register(&req)
	if err != nil{
		h.Logger.Error(fmt.Sprintf("Register request error: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Xatolik",
		})
		return 
	}
	c.JSON(http.StatusOK, resp)
}

func(h *hanlerimpl) Login(c *gin.Context){
	req := model.LoginReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil{
		h.Logger.Error(fmt.Sprintf("Ma'lumotlarni bodydan o'qishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Xato ma'lumotlar kiritildi",
		})
		return
	}

	user, err := h.Storage.UserService().CheckUser(req.Email)
	if err != nil{
		h.Logger.Error(fmt.Sprintf("CheckUser request error: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Bunday emailli foydalanuvchi mavjud emas",
		})
		return 
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil{
		h.Logger.Error("Password xato: %v", err)
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Password xato kiritildi.",
		})
		return 
	}

	access, err := token.GenerateAccessToken(user)
	if err != nil{
		h.Logger.Error(fmt.Sprintf("Access token generatsiya qilinmadi: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: err.Error(),
		})
		return 
	}	

	refresh, err := token.GenerateRefreshToken(user)
	if err != nil{
		h.Logger.Error(fmt.Sprintf("Refresh token generatsiya qilinmadi: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: err.Error(),
		})
		return 
	}	

	err = h.Storage.UserService().SaveToken(&model.SaveTokenReq{
		UserId: user.Id,
		RefreshToken: refresh,
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
	})
	if err != nil{
		h.Logger.Error(fmt.Sprintf("Tokenni saqlashda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Message: "Xatolik",
		})
		return
	}
	c.JSON(http.StatusOK, model.LoginResp{
		RefreshToken: refresh,
		AcessToken: access,
	})
}