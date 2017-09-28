package popcount_64shifts

//pc[i] jest liczebnoscia populacji i.
var pc [256]byte

//PopCountzwraca liczebnosc populacji (liczbe ustawionych bitow) dla x.
func PopCount(x uint64) int {
	var result int
	for i := 0; i < 64; i++ {
		result += int(x & 1)
		x >>= 1
	}
	return result

}
