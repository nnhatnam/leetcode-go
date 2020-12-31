package _108_Defanging_an_IP_Address

import "strings"

func defangIPaddr(address string) string {
	var b strings.Builder

	for _, v := range address {
		if v == '.' {
			b.WriteString("[.]")
		} else {
			b.WriteByte(byte(v))
		}
	}
	return b.String()
}
