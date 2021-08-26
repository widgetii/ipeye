package main

import (
	"fmt"
	"log"
	"time"

	"github.com/deepch/vdk/format/rtspv2"
)

func main() {
	url := "rtsp://10.216.128.64"
	//url := "rtsp://admin:admin@10.216.128.57"

	RTSPClient, err := rtspv2.Dial(rtspv2.RTSPClientOptions{URL: url, DisableAudio: false, DialTimeout: 3 * time.Second, ReadWriteTimeout: 3 * time.Second, Debug: true})
	if err != nil {
		log.Fatalln(err)
	}
	defer RTSPClient.Close()

	fmt.Printf("%#v", RTSPClient.CodecData)

	for {
		select {
		case signals := <-RTSPClient.Signals:
			switch signals {
			case rtspv2.SignalCodecUpdate:
				log.Println("SignalCodecUpdate")
			case rtspv2.SignalStreamRTPStop:
				log.Fatalln("ErrorStreamExitRtspDisconnect")
			}
		case packetAV := <-RTSPClient.OutgoingPacketQueue:
			fmt.Println(packetAV.Idx, len(packetAV.Data))
		}
	}
}
