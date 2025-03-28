package service

import (
	"WakeOnLan/internal/config"
	"bytes"
	"fmt"
	"log"
	"net"
)

type WakeService struct {
	config config.Config
	MAC    net.HardwareAddr
	packet []byte
}

func (s WakeService) WakeOnLAN() {
	bCastAddr := fmt.Sprintf("255.255.255.255:%d", s.config.Port+1)
	udpAddr, err := net.ResolveUDPAddr("udp4", bCastAddr)
	if err != nil {
		panic(fmt.Errorf("解析 UDP 地址失败: %w", err))
	}
	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		panic(fmt.Errorf("创建 UDP 连接失败: %w", err))
	}
	defer conn.Close()
	n, err := conn.Write(s.packet)
	if err != nil {
		panic(fmt.Errorf("发送 UDP 数据包失败: %w", err))
	}
	if n != len(s.packet) {
		panic(fmt.Errorf("发送 UDP 数据包不完整: %w", err))
	}
}

func CreateWakeService(config config.Config) *WakeService {
	hardwareAddr, err := net.ParseMAC(config.MAC)
	if err != nil {
		log.Fatal("无效的 MAC 地址:", config.MAC)
	}
	packet := bytes.NewBuffer(nil)
	packet.Write(bytes.Repeat([]byte{0xFF}, 6))
	for i := 0; i < 16; i++ {
		packet.Write(hardwareAddr)
	}
	return &WakeService{config, hardwareAddr, packet.Bytes()}
}
