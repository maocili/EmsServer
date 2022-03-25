package main

import (
	"EMSServer/EMS"
	"EMSServer/EMS/utils"
	"fmt"
	"net"
)

func main() {

	session := EMS.NewEMSSession("127.0.0.1:5555")
	go session.CatchInfo()
	res := session.GetIP()
	divceID := utils.RegNetDivceID(res)

	ip := net.ParseIP("10.0.2.100")
	subnet := net.ParseIP("255.255.255.0")
	gateway := net.ParseIP("10.0.0.1")

	issucce := session.SetIP(divceID, ip, subnet, gateway)
	if issucce {
		res = session.GetIP()
		fmt.Println(res)
	}
	for i := 0; i < 10; i++ {
		res = session.GetIP()
		fmt.Println(res)
	}

	session.Shutdown()

	for {
		select {}
	}

}
