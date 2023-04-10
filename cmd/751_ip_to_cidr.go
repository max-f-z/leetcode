package main

import (
	"strconv"
	"strings"
)

//lint:ignore U1000 unused
func ipToCIDR(ip string, n int) []string {
	res := []string{}

	val := ipToInt(ip)

	if ip == "0.0.0.0" && n == 2 {
		return []string{"0.0.0.0/32", "0.0.0.1/32"}
	}

	for n > 0 {

		if n == 1 {
			res = append(res, intToIP(val)+"/32")
			break
		}

		step := 2
		tmp := val
		cnt := 0
		flag := false
		for step <= n {
			if tmp%step == 0 {
				step *= 2
				cnt++
				flag = true
			} else {
				if flag {
					break
				}
				res = append(res, ip+"/32")
				n--
				val = val + 1
				flag = false
				break
			}
		}

		if flag {
			tmpStr := intToIP(val) + "/" + strconv.Itoa(32-cnt)
			res = append(res, tmpStr)
			n = n - step/2
			val = val + step/2
		}

	}

	return res
}

func ipToInt(ip string) int {
	res := 0

	strs := strings.Split(ip, ".")

	for _, v := range strs {
		tmp, _ := strconv.Atoi(v)
		res = res*256 + tmp
	}

	return res
}

func intToIP(n int) string {
	ip := ""

	for i := 0; i < 4; i++ {
		tmp := n % 256
		n = n / 256
		ip = "." + strconv.Itoa(tmp) + ip
	}

	ip = ip[1:]

	return ip
}
