package shift

func shift(s []int, n int) {
    reverse(s[:n])
    reverse(s[n:])
    reverse(s)
}


func reverse(s []int) {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j+1 {
        s[i], s[j] = s[j], s[i]
    }
}
