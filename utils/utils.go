package utils

import (
	"fmt"
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

func getBinary(num int64) string {
	return fmt.Sprintf("%064b", num)
}

func GetTypeOkBody(prefix string) map[string]interface{} {
	return map[string]interface{}{
		"type": GenerateTypeOk(prefix),
	}
}
