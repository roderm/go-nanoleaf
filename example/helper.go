package example

import (
	"net"
	"strconv"
	"strings"
)

func GetIp() net.IP {
	ip := make([]byte, 4)
	digits := strings.Split(IP, ".")
	for i, d := range digits {
		di, _ := strconv.Atoi(d)
		ip[i] = uint8(di)
	}
	return ip
}
