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
	//raddr:=&net.UDPAddr{IP:net.IPv4(byte(67),byte(215),byte(246),byte(10)),Port:6881}
	raddr:=&net.UDPAddr{IP:net.ParseIP("132.232.88.200"),Port:36556}
	conn, e := net.DialUDP("udp",nil,raddr)
	if e != nil {
		fmt.Printf("连接远程地址错误:%s",e.Error())
	}
	defer conn.Close()
	//i, e := conn.WriteToUDP([]byte("hello"), raddr)
	wp.Add(10)
	for a:=0;a<10;a++{
		go func(conn *net.UDPConn) {
			n, oobn, e := conn.WriteMsgUDP([]byte("d1：ad2：id20：abcdefghij0123456789e1：q4：ping1：t2：aa1：y1：qe"), nil, nil)
			//i, e := conn.WriteToUDP([]byte("d1：ad2：id20：abcdefghij0123456789e1：q4：ping1：t2：aa1：y1：qe"), raddr)
			if e != nil {
				fmt.Printf("发送数据错误:%s \r\n",e.Error())
				return

			}
			fmt.Println(n,oobn)
			wp.Done()
		}(conn)
	}
	wp.Wait()





}
