package config

import (
	"github.com/spf13/viper"
)

type Global struct {
	Max           int    //当前服务器主机允许的最大链接个数
	Pool          uint32 //业务工作Worker池的数量
	TaskLen       uint32
	MsgChanLen    uint32 //SendBuffMsg发送消息的缓冲最大长度
	MaxPacketSize uint32 //都需数据包的最大值
	Log           Zap
	TcpPort       int
	Ip            string
	HttpPort      string
	Name          string
}

// todo 自定义
type Zap struct {
	Mod        string
	Path       string
	LogType    string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

var (
	Conf Global
)

func Init() {

	var (
		v = viper.New()
	)

	v.SetConfigType("json")
	v.AddConfigPath(".")
	v.AddConfigPath("../")

	v.SetConfigName("config.json")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
		return
	}

	if err := v.Unmarshal(&Conf); err != nil {
		panic(err)
	}

}
