package EMS

import (
	"fmt"
	"net"
)

type EMSSession struct {
	Conn     net.Conn
	IPAddr   string
	Status   EMSSessionStatus
	Quit     chan int
	WaitSAC  chan int
	Response chan []byte
}

type EMSSessionStatus int

const (
	STATUS_INIT     EMSSessionStatus = 0
	STATUS_STANDBY  EMSSessionStatus = 1
	STATUS_RUNNING  EMSSessionStatus = 2
	STATUS_FINISH   EMSSessionStatus = 3
	STATUS_SHUTDOWN EMSSessionStatus = 4
)

func NewEMSSession(ipaddr string) (ems EMSSession) {

	server, err := net.Listen("tcp", ipaddr)
	if err != nil {
		fmt.Printf("EMS.NewEMSSession:%s", err)
	}
	fmt.Printf("EMS Session listen tcp: %s \n", ipaddr)

	conn, err := server.Accept()
	if err != nil {
		fmt.Printf("EMS.NewEMSSession:%s", err)
	}

	ems = newEms(conn, ipaddr)

	go func(server net.Listener) {
		for {
			select {
			case <-ems.Quit:
				fmt.Printf("Listen port %s will be closed", ipaddr)
				ems.Close(server)
			}
		}
	}(server)

	return ems
}

func newEms(conn net.Conn, ipaddr string) EMSSession {
	return EMSSession{
		Conn:     conn,
		IPAddr:   ipaddr,
		Status:   STATUS_INIT,
		Quit:     make(chan int, 1),
		WaitSAC:  make(chan int, 1),
		Response: nil,
	}
}

func (session *EMSSession) Close(server net.Listener) {
	session.Conn.Close()
	server.Close()
}
