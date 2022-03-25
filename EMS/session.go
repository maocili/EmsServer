package EMS

import (
	"EMSServer/EMS/utils"
	"fmt"
	"log"
)

func (session *EMSSession) CatchInfo() {

	b := make([]byte, 1024)
	var response []byte

	fmt.Println(session.Conn.RemoteAddr().String())

	for {
		n, err := session.Conn.Read(b)
		if err != nil {
			log.Printf("[%s]: Connection has been closed by remote peer, %d bytes has been received\n", session.Conn.RemoteAddr(), n)
			log.Printf("[%s]: Local peer has been stopped, %d bytes\n", session.Conn.RemoteAddr(), n)
			session.Quit <- 1
			return
		}

		for i := 0; i < n; i++ {
			if b[i] == '\n' || b[i] == '>' {
				session.statusSwitch(response)
				response = []byte{}
			} else {
				response = append(response, b[i])
			}
		}

	}
}

func (session *EMSSession) statusSwitch(response []byte) {

	fmt.Printf("Status: %d | %s \n", session.Status, response)

	switch session.Status {
	case STATUS_INIT:
		session.waitSAC(response)
	case STATUS_RUNNING:
		session.waitResponse(response)
	case STATUS_SHUTDOWN:
		session.waitShutdown(response)
	}
}
func (session *EMSSession) waitShutdown(response []byte) {

	if utils.MatchShutdown(response) {
		session.Quit <- 1
	}
}
func (session *EMSSession) waitResponse(response []byte) {
	if utils.MatchSAC(response) {
		session.Status = STATUS_FINISH
		session.Response <- []byte{03}
		session.WaitSAC <- 0

		return
	}
	session.Response <- response
}

func (session *EMSSession) waitSAC(response []byte) {
	if utils.MatchSAC(response) {
		session.WaitSAC <- 0
		switch session.Status {
		case STATUS_INIT:
			session.Status = STATUS_STANDBY
		case STATUS_RUNNING:
			session.Status = STATUS_FINISH
		}
	}
}
