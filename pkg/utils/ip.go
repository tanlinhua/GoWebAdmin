package utils

import (
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"regexp"
)

const (
	v4Seg   = "(?:[0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])"
	v4Str   = "(" + v4Seg + "[.]){3}" + v4Seg
	ipv4Reg = "^" + v4Str + "$"
	v4URL   = "https://api.ipify.org"
)

const (
	v6Seg   = "(?:[0-9a-fA-F]{1,4})"
	ipv6Reg = "^(" +
		"(?:" + v6Seg + ":){7}(?:" + v6Seg + "|:)|" +
		"(?:" + v6Seg + ":){6}(?:" + v4Str + "|:" + v6Seg + "|:)|" +
		"(?:" + v6Seg + ":){5}(?::" + v4Str + "|(:" + v6Seg + "){1,2}|:)|" +
		"(?:" + v6Seg + ":){4}(?:(:" + v6Seg + "){0,1}:" + v4Str + "|(:" + v6Seg + "){1,3}|:)|" +
		"(?:" + v6Seg + ":){3}(?:(:" + v6Seg + "){0,2}:" + v4Str + "|(:" + v6Seg + "){1,4}|:)|" +
		"(?:" + v6Seg + ":){2}(?:(:" + v6Seg + "){0,3}:" + v4Str + "|(:" + v6Seg + "){1,5}|:)|" +
		"(?:" + v6Seg + ":){1}(?:(:" + v6Seg + "){0,4}:" + v4Str + "|(:" + v6Seg + "){1,6}|:)|" +
		"(?::((?::" + v6Seg + "){0,5}:" + v4Str + "|(?::" + v6Seg + "){1,7}|:))" +
		")(%[0-9a-zA-Z-.:]{1,})?$"
	v6URL = "https://api64.ipify.org"
)

func IsIPv4(s string) bool {
	ok, err := regexp.MatchString(ipv4Reg, s)
	if err != nil {
		return false
	}

	return ok
}

func IsIPv6(s string) bool {
	ok, err := regexp.MatchString(ipv6Reg, s)
	if err != nil {
		return false
	}

	return ok
}

func GetIPV4() (string, error) {
	resp, err := http.Get(v4URL)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body := resp.Body
	txt, err := ioutil.ReadAll(body)

	if err != nil {
		return "", err
	}

	ip := string(txt)

	if !IsIPv4(ip) {
		return ip, errors.New(ip + ": not ipv4")
	}

	return ip, nil
}

func GetIPV6() (string, error) {
	resp, err := http.Get(v6URL)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body := resp.Body
	txt, err := ioutil.ReadAll(body)

	if err != nil {
		return "", err
	}

	ip := string(txt)

	if !IsIPv6(ip) {
		return ip, errors.New(ip + ": not ipv6")
	}

	return ip, nil
}

// GetServerIP 获取服务器外网IP地址,ip.String()
func GetServerIP() (net.IP, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return nil, err
		}
		for _, addr := range addrs {
			ip := getIpFromAddr(addr)
			if ip == nil {
				continue
			}
			return ip, nil
		}
	}
	return nil, errors.New("connected to the network?")
}

func getIpFromAddr(addr net.Addr) net.IP {
	var ip net.IP
	switch v := addr.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	}
	if ip == nil || ip.IsLoopback() {
		return nil
	}
	ip = ip.To4()
	if ip == nil {
		return nil // not an ipv4 address
	}
	return ip
}
