package util

func IsBitSet(value uint32, n int) bool {
	return ((value >> n) & 1) != 0
}
