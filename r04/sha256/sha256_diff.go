package main

import (
    "fmt"
    "crypto/sha256"
    "os"
    "../../r02/popcount"
    "encoding/binary"
)

func main() {
    if len(os.Args) < 2 || len(os.Args) > 3 {
        fmt.Printf("\nusage: sha256_diff arg1 arg2\n Count bit difference between 2 sha256 hashes.\n")
    }else {
        args := os.Args[1:]
        c1 := sha256.Sum256([]byte(args[0]))
        c2 := sha256.Sum256([]byte(args[1]))
        fmt.Printf("arg1: %s sha256: %x\n", args[0], c1)
        fmt.Printf("arg2: %s sha256: %x\n", args[1], c2)
        x := xor(c1, c2)
        var result int
        var ui uint64
        for i := 0; i < 4; i++ {
            ui = binary.LittleEndian.Uint64(x[8*i:(8*(i+1))])
            result += popcount.PopCount(ui)
        }
        fmt.Printf("\nresult count: %d\n", result)
    }
}

func xor(c1, c2 [32]byte) [32]byte {
    var r [32]byte
    for i := range r {
        r[i] = c1[i]^c2[i]
    }
    return r
}
