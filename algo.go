package unique_ipv4

import (
	"strconv"
	"strings"
)

const MaxUint32 = ^uint32(0)
const IPNumbersLen = 4

func ConvertIPv4SttringToUint32(ipv4 string) uint32 {
	stringNumbers := strings.Split(ipv4, ".")
	var result uint32 = 0
	for i := range IPNumbersLen {
		number, _ := strconv.Atoi(stringNumbers[i])
		result |= uint32(number) << ((IPNumbersLen - i - 1) * 8)
	}
	return result
}

func GetTotalUniqueIPv4Addresses(addresses []string) uint32 {
	bitSet := make([]bool, MaxUint32, MaxUint32)
	var result uint32 = 0
	for _, address := range addresses {
		n := ConvertIPv4SttringToUint32(address)
		if !bitSet[n] {
			bitSet[n] = true
			result++
		}
	}
	return result
}
