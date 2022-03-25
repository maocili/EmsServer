package utils

import "bytes"

//Match TCP info Flag
const (
	FLAGSAC      = "SAC"
	FLAGSHUTDOWN = "计算机正在关闭"
)

func MatchShutdown(response []byte) bool {
	return matchInfo(response, []byte(FLAGSHUTDOWN))

}

func MatchSAC(response []byte) bool {
	return matchInfo(response, []byte(FLAGSAC))
}

func matchInfo(info []byte, flag []byte) bool {
	return bytes.Compare(info, flag) == 0
}
