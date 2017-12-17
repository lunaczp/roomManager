package config

import (
	"time"
)

type BroadcastingStationOption struct {
	Host string
	Port int64
	Path string
}

var (
	UseBoradcasting         = false                      //是否使用广播站
	RoomCleanTimer          = time.Minute * 5            //房间清理时间间隔
	MaxLengthOfRows         = int64(1024)                //房间内单行最大节点数
	MaxMessageChannelLength = int64(128)                 //消息通道最大长度
	broadcastingOption      = BroadcastingStationOption{ //广播站连接参数
		Host: "localhost",
		Port: 9981,
		Path: "broadcasting",
	}
)

//设定广播站连接方式
func SetBroadcastingStation(opt BroadcastingStationOption) {
	broadcastingOption = opt
	UseBoradcasting = true
}

func GetBroadcastingStation() BroadcastingStationOption {
	return broadcastingOption
}

func SetRoomCleanTimer(t time.Duration) {
	RoomCleanTimer = t
}

func SetMaxLengthOfRows(length int64) bool {
	if length <= 0 {
		return false
	}
	MaxLengthOfRows = length
	return true
}