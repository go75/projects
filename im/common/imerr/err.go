package imerr

import "errors"

// token
var ExpiredTokenErr = errors.New("token过期")
var InvaildTokenErr = errors.New("token无效")

// websocket
var AlreadyExistConnErr = errors.New("websocket连接已存在")