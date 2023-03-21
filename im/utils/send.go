package utils

import (
	"encoding/json"
	"errors"
	"im/global"
	"im/ws/entity"
	"im/ws/manager"

	"github.com/gorilla/websocket"
)

func Send(connID, id, option, processId uint, payload any) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	obj := entity.Obj {
		ID: id,
		Type: option,
		ProcessId: processId,
		Payload: string(data),
	}
	data, err = json.Marshal(obj)
	if err != nil {
		return err
	}
	conn := manager.Get(connID)
	if conn == nil {
		return errors.New("连接不在线")
	}
	err = conn.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		return err
	}
	return nil
}

func Publish(channel string, id, option, processId uint, payload any) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	obj := entity.Obj{
		ID: id,
		Type: option,
		ProcessId: processId,
		Payload: string(data),
	}
	data, err = json.Marshal(obj)
	if err != nil {
		return err
	}
	global.Rd.Publish(channel, string(data))
	return nil
}