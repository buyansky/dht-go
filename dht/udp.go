package dht

import (
	"fmt"
	"net"
	"sync"
)

var wp sync.WaitGroup

type Udp struct {
	Laddr *net.UDPAddr
}

func (u *Udp)GetUdpConn()  {

}


//发送udp 数据
func (u *Udp)SetUdpData(data []byte)  {

	raddr:=&net.UDPAddr{IP:net.ParseIP("67.215.246.10"),Port:6881}
	conn, e := net.DialUDP("udp",nil,raddr)
	if e != nil {
		fmt.Printf("连接远程地址错误:%s",e.Error())
		return 
	}
	defer conn.Close()


	wp.Add(10)
	for a:=0;a<10;a++{
		go func(conn *net.UDPConn) {

			i, err := conn.Write(data)
			if err != nil {
				fmt.Printf("发送数据错误:%s \r\n",err.Error())
				return

			}
			fmt.Printf("接收到的数据:%d\r\n",i)
			wp.Done()
		}(conn)
	}
	wp.Wait()





}
