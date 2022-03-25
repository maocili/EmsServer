package EMS

import (
	"fmt"
	"io"
	"net"
	"strings"
)

func (session *EMSSession) GetIP() string {

	result, err := session.sendCommand("i")

	if err != nil {
		fmt.Errorf("EMS.GetIP:", err)
		return ""
	}

	return result

}

func (session *EMSSession) SetIP(deviceID string, ip net.IP, subnet net.IP, gateway net.IP) bool {

	cmd := fmt.Sprintf("i %s %s %s %s", deviceID, ip, subnet, gateway)
	result, err := session.sendCommand(cmd)
	if err != nil {
		fmt.Printf("EMS.SetIP:%s", err)
	}

	return strings.Contains(result, "未能设置")
}

func (session *EMSSession) Shutdown() {
	_, err := session.sendCommand("shutdown")
	if err != nil {
		fmt.Println("ems.Shutdown:", err)
	}

	session.Status = STATUS_SHUTDOWN

}

func (session *EMSSession) sendCommand(cmd string) (result string, err error) {

	<-session.WaitSAC
	if session.Status == STATUS_STANDBY {
		session.Status = STATUS_RUNNING
		session.Response = make(chan []byte)
	}
	_, err = io.Copy(session.Conn, strings.NewReader(fmt.Sprintf("%s\n", cmd)))

	if err != nil {
		return "", err
	}

	for session.Status == STATUS_RUNNING {
		result += fmt.Sprintf("%s", <-session.Response)
	}

	session.Status = STATUS_STANDBY

	return result, err
}
