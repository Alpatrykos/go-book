package main

import (
    "fmt"
    "os"
    "strconv"
    "./tempconv"
    "./distconv"
)

func main() {
    for _, arg := range os.Args[1:] {
        t, err := strconv.ParseFloat(arg, 64)
        if err != nil {
            fmt.Fprintf(os.Stderr, "%v\n", err)
            os.Exit(1)
        }
        f := tempconv.Fahrenheit(t)
        c := tempconv.Celsius(t)
        fmt.Printf("%s = %s, %s = %s\n",
            f, tempconv.FToC(f), c, tempconv.CToF(c))
        m := distconv.Meter(t)
        ft := distconv.Foot(t)
        fmt.Printf("%s = %s, %s = %s\n",
            m, distconv.MToFt(m), ft, distconv.FtToM(ft))
    }
}
