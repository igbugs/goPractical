package main

import (
	"bufio"
	"day15/tcp_chat/protocol"
	"fmt"
	"golang.org/x/net/context"
	"logging"
	"math/rand"
	"net"
	"os"
	"time"
)

var (
	roomList *protocol.AllRoomList
	user     *UserInfo
)

func init() {
	//每次启动一个client 的客户端，分配一个唯一的用户ID
	rand.Seed(time.Now().UnixNano())
	userId := rand.Int63()
	user = &UserInfo{
		UserID:   uint64(userId),
		UserName: fmt.Sprintf("user%d", userId),
	}

	logging.Error("generate default user_info, user_id:%d, user_name:%s", user.UserID, user.UserName)
}

func main() {
	conn, err := net.Dial("tcp", "localhost:50000")
	if err != nil {
		logging.Error("dialing error, err:%v", err)
		return
	}
	defer conn.Close()

	//1. 拉去房间列表
	logging.Debug("start get room list")
	roomList, err = getRoomList(conn)
	if err != nil {
		logging.Error("get room list failed, err:%v", err)
		return
	}

	for {
		showRoomList(roomList)
		var roomId uint64
		fmt.Scanf("%d\n", &roomId)
		roomInfo, err := enterRoom(conn, roomId)
		if err != nil {
			logging.Error("enter room failed, err:%v", err)
			continue
		}

		processRoomMessage(conn, roomInfo)
	}
}

func processRoomMessage(conn net.Conn, roomInfo *protocol.Room) {
	fmt.Printf("enter room success\n")
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
	}()

	go recvMessage(ctx, conn)
	for {
		reader := bufio.NewReader(os.Stdin)
		msg, err := reader.ReadString('\n')
		if err != nil {
			logging.Error("read string failed, err:%v", err)
			continue
		}

		err = sendMessage(conn, msg, roomInfo)
		if err != nil {
			logging.Error("send msg failed, err:%v", err)
			continue
		}

	}
}

func recvMessage(ctx context.Context, conn net.Conn) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		typ, result, err := protocol.UnPack(conn)
		if err != nil {
			logging.Error("unpack message failed, err:%v", err)
			return
		}

		switch typ {
		case protocol.BROADCAST_USER_ENTER_ROOM:
			broadcastEnterRoom, ok := result.(*protocol.BroadcastEnterRoom)
			if !ok {
				break
			}
			fmt.Printf("username:%s enter room\n", broadcastEnterRoom.EnterUserName)
		case protocol.USER_RECV_TEXT:
			recvMsg, ok := result.(*protocol.UserRecvText)
			if !ok {
				break
			}

			fmt.Printf("%s:\n", recvMsg.AuthorUserName)
			fmt.Printf("  %s", recvMsg.Content)
		}
	}
}

func sendMessage(conn net.Conn, msg string, roomInfo *protocol.Room) (err error) {
	userSendMsg := &protocol.UserSendText{
		RoomID:   roomInfo.RoomID,
		UserID:   user.UserID,
		UserName: user.UserName,
		Content:  msg,
	}

	data, err := protocol.Pack(protocol.USER_SEND_TEXT, userSendMsg)
	if err != nil {
		logging.Error("sendMessage pack failed, err:%v", err)
		return
	}

	_, err = conn.Write(data)
	if err != nil {
		logging.Error("write message to server failed, err:%v", err)
		return
	}

	return
}
