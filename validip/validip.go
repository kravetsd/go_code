package kata

import (
	"fmt"
	"net"
	"strings"
)

func Is_valid_ip(ip string) bool {

	ipv4 := net.ParseIP(ip)

	if ipv6 := ipv4.To16(); ipv6 != nil {
		if j := ipv6.To4(); j != nil {
			sbst := strings.Split(ip, ".")
			for _, v := range sbst {
				fmt.Println(sbst)
				fmt.Println(v[0])
				if len(v) > 1 && v[0] == 48 {
					return false
				}
			}
			return true
		}
		return false
	}
	return false
}
