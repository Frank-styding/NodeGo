package utils

func GetBit(value int, bitIndex int) int {
	return (value >> bitIndex) & 1
}

func SetBit(value *int, bitIndex int, bitValue int) {
	if bitValue == 1 {
		*value |= (1 << bitIndex)
	} else {
		*value &= ^(1 << bitIndex)
	}
}