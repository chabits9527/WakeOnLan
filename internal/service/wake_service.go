package service

import (
	"bytes"
	"fmt"
	"log"
	"net"
)

type WakeService struct {
	MAC    net.HardwareAddr
	packet []byte
}

func (s WakeService) WakeOnLAN() {
	fmt.Printf("发送唤醒信号到 %v", s.MAC)
}

func CreateWakeService(MAC string) *WakeService {
	hardwareAddr, err := net.ParseMAC(MAC)
	if err != nil {
		log.Fatal("无效的 MAC 地址:", MAC)
	}
	packet := bytes.NewBuffer(nil)
	packet.Write(bytes.Repeat([]byte{0xFF}, 6))
	for i := 0; i < 16; i++ {
		packet.Write(hardwareAddr)
	}
	return &WakeService{hardwareAddr, packet.Bytes()}
}
