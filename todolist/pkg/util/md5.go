package util

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5Encode(data string) string {	
	h := md5.New()
	h.Write([]byte(data))
	tempStr := h.Sum(nil)
	return hex.EncodeToString(tempStr)
}

func MakePwd(rawPwd string) string {
	return Md5Encode(rawPwd)
}

func CheckPwd(rawPwd, pwd string) bool {
	return MakePwd(rawPwd) == pwd
}
