package ip

import (
	`fmt`
	`math/rand`
	`time`
)

/*
	Created by aomerk at 2021-11-21 for project strixeye
*/

/*
	Ip represents an IPv4 address.
*/

// global constants for file
const ()

// global variables (not cool) for this file
var ()

type IPv4 [4]uint8

func NewRandom() IPv4 {
	ip := IPv4{}
	max := uint8(255)
	rand.Seed(time.Now().UnixNano())
	ip[3] = uint8(rand.Intn(int(max)))
	ip[2] = uint8(rand.Intn(int(max)))
	ip[1] = uint8(rand.Intn(int(max)))
	ip[0] = uint8(rand.Intn(int(max)))
	return ip
}

func RandomPort() string {
	max := uint16(65535)
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%d", rand.Intn(int(max)))
}

func Parse(ipString string) IPv4 {
	ip := IPv4{}
	_, err := fmt.Sscanf(ipString, "%d.%d.%d.%d", &ip[0], &ip[1], &ip[2], &ip[3])
	if err != nil {
		panic(err)
	}
	return ip
}

func (d IPv4) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", d[3], d[2], d[1], d[0])
}
