package utils

import (
	"strconv"
)

func GenerateTypeOk(prefix string) string {
	return prefix + "_ok"
}

func GetNodeIDInt(nid string) int {
	id, _ := strconv.Atoi(nid[1:])
	return id
}

func GetUniqueID(nodeID, num int64) int64 {
	return (nodeID << 48) | num
}
