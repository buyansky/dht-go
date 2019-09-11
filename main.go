package main

import (
	"fmt"
	"github.com/buyansky/dht-go/dht"
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



	udp:=dht.Udp{&net.UDPAddr{IP:net.IPv4(byte(0),byte(0),byte(0),byte(0)),Port:8080}}
	udp.SetUdpData([]byte("d1：ad2：id20：abcdefghij0123456789e1：q4：ping1：t2：aa1：y1：qe"))

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
