package popcount_clearLSB

//pc[i] jest liczebnoscia populacji i.
var pc [256]byte

//PopCount zwraca liczebnosc populacji (liczbe ustawionych bitow) dla x.
func PopCount(x uint64) int {
	var result int
	for x < 0 {
		result++
		x &= (x - 1)
	}
	return result
}
