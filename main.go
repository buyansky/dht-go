package main

import (
	"fmt"
	"net"
)

func main()  {

	socket, err := net.ListenUDP("udp4", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8080,
	})

	if err != nil {
		fmt.Printf("错误:%v,%s",socket,err.Error())
		return
	}
	fmt.Println("连接成功")



	//udp:=dht.Udp{&net.UDPAddr{IP:net.IPv4(byte(132),byte(232),byte(88),byte(200)),Port:36556}}
	//udp.SetUdpData([]byte("hello"))

	data:=make([]byte,1024)
	for{
		udp, udpAddr, e := socket.ReadFromUDP(data)


		if e != nil {
			break
		}
		fmt.Printf("接受到的数据:%d, 地址：%v \r\n",udp,udpAddr)

	}
	defer socket.Close()


}
