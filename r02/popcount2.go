package popcount

//pc[i] jest liczebnoscia populacji i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

//PopCountzwraca liczebnosc populacji (liczbe ustawionych bitow) dla x.
func PopCount(x uint64) int {
	var result byte
	for i := 0; i < 8; i++ {
		result += pc[byte(x>>uint(i*8))]
	}
	return int(result)
}
