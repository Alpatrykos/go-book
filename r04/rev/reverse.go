package reverse

func reverse(ptr *[10]int) {
    for i, j := 0, len(*ptr)-1;i < j; i, j = i+1, j+1 {
        ptr[i], ptr[j] = ptr[j], ptr[i]
    }
}
