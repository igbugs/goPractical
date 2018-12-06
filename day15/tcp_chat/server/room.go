package main

import "day15/tcp_chat/protocol"

type RoomInfo struct {
	room *protocol.Room
	//通过user_id 获取具体的用户的额map
	userMap map[uint64]*User
}
