package main

import (
	"fmt"
	"github.com/buyansky/dht-go/dht"
	"net"
)

func main()  {

	uip:=net.ParseIP("132.232.88.200")
	udpaddr:=net.UDPAddr{IP:uip,Port:36556}
	listener,e:= net.ListenUDP("udp",&udpaddr)

	if e != nil {
		fmt.Printf("错误:%v,%s",listener,e.Error())
	}
	fmt.Println("连接成功")
	//defer listener.Close()


	udp:=dht.Udp{&net.UDPAddr{IP:net.IPv4(byte(132),byte(232),byte(88),byte(200)),Port:36556}}
	udp.SetUdpData([]byte("hello"))

	data:=make([]byte,1024)
	for{
		udp, udpAddr, e := listener.ReadFromUDP(data)


		if e != nil {
			break
		}
		fmt.Printf("接受到的数据:%d, 地址：%v \r\n",udp,udpAddr)

	}



}
