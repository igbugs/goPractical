package controller

import (
	"golang.org/x/net/websocket"
	"web_chat/models"
)

type UserInfo struct {
	User *models.User
	Conn *websocket.Conn
	RoomInfo *RoomInfo
}

