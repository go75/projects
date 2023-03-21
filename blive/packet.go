package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"compress/zlib"
	"github.com/gorilla/websocket"
)

type Packet struct {
	PacketLen  int
	HeaderLen  int
	Version    int
	Operation  int
	SequenceID int

	Body []PacketBody
}

type PacketBody struct {
	Cmd     string                 `json:"cmd"`
	Data    map[string]interface{} `json:"data"`
	MsgSelf string                 `json:"msg_self"`
	Info    []interface{}          `json:"info"`
}

func (c *BiliClient) SendPacket(packetLen uint32, headerLen uint16, version uint16, operation uint32, sequenceID uint32, body []byte) error {
	if packetLen == 0 {
		packetLen = uint32(len(body) + 16)
	}

	header := new(bytes.Buffer)

	var data = []any{
		packetLen,
		headerLen,
		version,
		operation,
		sequenceID,
	}

	for _, v := range data {
		err := binary.Write(header, binary.BigEndian, v)
		if err != nil {
			return err
		}
	}

	socketData := append(header.Bytes(), body...)

	err := c.Conn.WriteMessage(websocket.TextMessage, socketData)

	return err
}

func (c *BiliClient) ReceiveMessages() {

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			fmt.Println("client ReceiveMessages err: ", err)
		}

		packet, err := c.Decode(message)
		if err != nil {
			fmt.Println("client Decode err: ", err)
		}

		for _, body := range packet.Body {
			c.Ch <- body
		}
	}

}

func (c *BiliClient) Decode(blob []byte) (Packet, error) {
	result := Packet{
		PacketLen:  int(binary.BigEndian.Uint32(blob[0:4])),
		HeaderLen:  int(binary.BigEndian.Uint16(blob[4:6])),
		Version:    int(binary.BigEndian.Uint16(blob[6:8])),
		Operation:  int(binary.BigEndian.Uint32(blob[8:12])),
		SequenceID: int(binary.BigEndian.Uint32(blob[12:16])),
		Body:       make([]PacketBody, 0),
	}

	if result.Operation == 5 {
		offset := 0

		for offset < len(blob) {
			packetLen := int(binary.BigEndian.Uint32(blob[offset : offset+4]))

			if result.Version == 2 {
				// zipped packet
				data := blob[offset+result.HeaderLen : offset+packetLen]

				r, err := zlib.NewReader(bytes.NewReader(data))
				if err != nil {
					return Packet{}, err
				}
				defer r.Close()

				var newBlob []byte
				newBlob, err = io.ReadAll(r)
				if err != nil {
					return Packet{}, err
				}

				var child Packet
				child, err = c.Decode(newBlob)
				if err != nil {
					return Packet{}, err
				}

				result.Body = append(result.Body, child.Body...)

			} else {
				data := blob[offset+result.HeaderLen : offset+packetLen]

				var body PacketBody
				if err := json.Unmarshal(data, &body); err != nil {
					return Packet{}, err
				}
				result.Body = append(result.Body, body)
			}
			offset += packetLen
		}
	}
	return result, nil
}

func (bc *BiliClient) ParseMessages() {
	for msg := range bc.Ch {
		switch msg.Cmd {
		case "COMBO_SEND":
			fmt.Printf("\u001B[1;36m[ %s 送给 %s %b 个 %s ]\u001B[0m\n", msg.Data["uname"].(string), msg.Data["r_uname"].(string), int(msg.Data["combo_num"].(float64)), msg.Data["gift_name"].(string))
		case "DANMU_MSG":
			fmt.Printf("\u001B[1;35m[ %s: %s ]\u001B[0m\n", msg.Info[2].([]interface{})[1].(string), msg.Info[1].(string))
		case "INTERACT_WORD":
			fmt.Printf("\u001B[1;34m[ %s 进入直播间 ]\u001B[0m\n", msg.Data["uname"].(string))
		case "SEND_GIFT":
			fmt.Printf("\u001B[1;33m[ %s 投喂 %b 个 %s ]\u001B[0m\n", msg.Data["uname"], int(msg.Data["num"].(float64)), msg.Data["giftName"].(string))
		case "NOTICE_MSG":
			fmt.Printf("\u001B[1;32m[ %s ]\u001B[0m\n", msg.MsgSelf)
		default:
			continue
		}
	}
}