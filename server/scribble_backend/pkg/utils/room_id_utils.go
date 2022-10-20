package utils

import (
	"math/rand"
	"time"
)

type RoomIDUtils struct {
	UsedKeys map[string]bool
}

const TMP = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

var roomUtils = RoomIDUtils{UsedKeys: make(map[string]bool)}

// key will be 4 character long
func CreateNewKey(roomUtils *RoomIDUtils) string {
	var s1 = rand.NewSource(time.Now().UnixNano())
	var r1 = rand.New(s1)
	key := ""
	for i := 0; i < 4; i++ {
		idx := r1.Intn(62)
		key += string(TMP[idx])
	}
	if _, ok := roomUtils.UsedKeys[key]; !ok {
		roomUtils.UsedKeys[key] = true
		return key
	}
	return CreateNewKey(roomUtils)
}

func GetKey() string {
	return CreateNewKey(&roomUtils)
}
