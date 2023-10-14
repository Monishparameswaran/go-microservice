package ContainerInfo

import (
	"log"
	"net"
	"os"
)

func Hostname() string {
	data, _ := os.Hostname()
	return data
}
func IP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
