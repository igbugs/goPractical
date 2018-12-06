package protocol

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"logging"
	"net"
)

const (
	USER_ENTER_ROOM           = 1001
	USER_LEAVE_ROOM           = 1002
	USER_SEND_TEXT            = 1003
	USER_RECV_TEXT            = 1004
	ALL_ROOM_LIST             = 1005
	GET_ROOM_LIST             = 1006
	USER_ENTER_ROOM_RESP      = 1007
	BROADCAST_USER_ENTER_ROOM = 1008
)

type Proto struct {
	Length uint32
	Type   uint16
}

type Room struct {
	RoomID  uint64
	Name    string
	RoomCap uint32
	Desc    string
	Online  uint32
}

type UserEnterRoom struct {
	RoomID   uint64
	UserID   uint64
	UserName string
}

type UserEnterRoomResp struct {
	Code     int
	RoomInfo *Room
}

type UserLeaveRoom struct {
	RoomID uint64
	UserID uint64
	UserName string
}

type UserSendText struct {
	RoomID   uint64
	UserID   uint64
	UserName string
	Content  string
}

type UserRecvText struct {
	Code           int
	RoomID         uint64
	AuthorUserID   uint64
	AuthorUserName string
	Content        string
}

type AllRoomList struct {
	RoomList []*Room
}

type GetRoomList struct {
	UserId uint64
}

type BroadcastEnterRoom struct {
	EnterUserID   uint64
	EnterUserName string
	RoomInfo      *Room
}

func Pack(typ uint16, body interface{}) (data []byte, err error) {
	bodyJson, err := json.Marshal(body)
	if err != nil {
		logging.Error("marshal json failed, body:%v, err:%v", body, err)
		return
	}

	var proto = &Proto{
		Length: uint32(len(bodyJson)),
		Type:   typ,
	}

	var buff bytes.Buffer
	err = binary.Write(&buff, binary.BigEndian, proto.Length)
	err = binary.Write(&buff, binary.BigEndian, proto.Type)
	err = binary.Write(&buff, binary.BigEndian, bodyJson)

	data = buff.Bytes()
	return
}

func UnPack(conn net.Conn) (typ uint16, result interface{}, err error) {
	//1.读取包的长度和包的类型
	/**
	 * type Proto struct {
	 * 	Length uint32   // 4个字节
	 *	Type   uint16   // 2个字节
	 * }
	 * buf 的字节slice长度为6 是因为 Proto 结构体定义的长度为 6 个字节
	 */
	buf := make([]byte, 6)
	_, err = conn.Read(buf)
	if err != io.EOF && err != nil {
		logging.Error("read package header err:%v", err)
		return
	}

	var proto Proto
	buffer := bytes.NewBuffer(buf)
	err = binary.Read(buffer, binary.BigEndian, &proto.Length)
	err = binary.Read(buffer, binary.BigEndian, &proto.Type)

	//2. 读取包体
	typ = proto.Type
	bodyBuf := make([]byte, proto.Length)
	_, err = conn.Read(bodyBuf)
	if err != nil {
		logging.Error("read from network failed, err:%v", err)
		return
	}

	//3. 根据消息的类型号， 反序列化对应的消息类型的结构体
	var tmpBody interface{}
	switch proto.Type {
	case GET_ROOM_LIST:
		tmpBody = &GetRoomList{}
	case ALL_ROOM_LIST:
		tmpBody = &AllRoomList{}
	case USER_ENTER_ROOM:
		tmpBody = &UserEnterRoom{}
	case USER_ENTER_ROOM_RESP:
		tmpBody = &UserEnterRoomResp{}
	case USER_LEAVE_ROOM:
		tmpBody = &UserLeaveRoom{}
	case USER_SEND_TEXT:
		tmpBody = &UserSendText{}
	case USER_RECV_TEXT:
		tmpBody = &UserRecvText{}
	case BROADCAST_USER_ENTER_ROOM:
		tmpBody = &BroadcastEnterRoom{}
	default:
		err = fmt.Errorf("unsupport message type:%d", proto.Type)
		logging.Error("unsupport message type:%d, data:%v", proto.Type, string(bodyBuf))
		return
	}

	err = json.Unmarshal(bodyBuf, tmpBody)
	if err != nil {
		logging.Error("unmarshal message body failed, err:%v", err)
		return
	}

	result = tmpBody
	logging.Debug("get message succ, type:%d. body length:%d, data:%#v",
		typ, proto.Length, result)
	return
}
