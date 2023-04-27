package utils

import (
	"errors"
	"net"
	"strconv"
	"strings"
)

// ipA:192.168.1.1
// ipB:192.168.1.0/24
// ipC:192.168.0.1-192.168.1.1
func GetIpList(ipFormat string) *[]string {

	ipFormatList := strings.Split(ipFormat, ";")

	rs := &[]string{}
	for _, v := range ipFormatList {
		v = strings.TrimSpace(v)
		if strings.Contains(v, "/") {
			b, err := ipHaveB(v)
			if err == nil {
				*rs = append(*rs, *b...)
			}
		}
		if strings.Contains(v, "-") {
			c, err := ipHaveC(v)
			if err == nil {
				*rs = append(*rs, *c...)
			}
		}
		_, err := ipParse(v)
		if err == nil {
			*rs = append(*rs, v)
		}
	}
	return rs
}

func ipParse(ipStr string) (uint32, error) {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return uint32(0x00000000), errors.New("ParseIP error")
	}
	ipInt := uint32(0x00000000)
	ipInt = ipInt | uint32(ip[12])<<24
	ipInt = ipInt | uint32(ip[13])<<16
	ipInt = ipInt | uint32(ip[14])<<8
	ipInt = ipInt | uint32(ip[15])
	return ipInt, nil
}

func ipHaveB(ipStr string) (*[]string, error) {
	ips := strings.Split(ipStr, "/")
	at, err := strconv.Atoi(ips[1])
	if err != nil {
		return nil, err
	}
	ipInt, err := ipParse(ips[0])
	if err != nil {
		return nil, err
	}
	start := ipInt & (uint32(0xffffffff) << (32 - at))
	end := ipInt | (uint32(0xffffffff) >> at)
	return ipGetList(start, end), nil
}

func ipHaveC(ipStr string) (*[]string, error) {
	ips := strings.Split(ipStr, "-")
	start, err := ipParse(ips[0])
	if err != nil {
		return nil, err
	}
	end, err := ipParse(ips[1])
	if err != nil {
		return nil, err
	}
	if start > end {
		return nil, errors.New("ip : start > end")
	} else {
		return ipGetList(start, end), nil
	}
}

func ipGetList(start, end uint32) *[]string {
	ls := []string{}
	for i := start; i <= end; i++ {
		a := strconv.Itoa(int((i & uint32(0xff000000)) >> 24))
		b := strconv.Itoa(int((i & uint32(0x00ff0000)) >> 16))
		c := strconv.Itoa(int((i & uint32(0x0000ff00)) >> 8))
		d := strconv.Itoa(int(i & uint32(0x000000ff)))
		ls = append(ls, a+"."+b+"."+c+"."+d)
	}
	return &ls
}
