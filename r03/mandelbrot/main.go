package main

import (
    "image"
    "image/color"
    "image/png"
    "math/cmplx"
    "os"
    "math"
)

func main() {
    const (
        xmin, ymin, xmax, ymax = -2, -2, +2, +2
        width, height = 1024, 1024
    )

    img := image.NewRGBA(image.Rect(0, 0, width, height))
    for py := 0; py < height; py++ {
        y := float64(py)/height*(ymax-ymin) + ymin
        for px := 0; px < width; px++ {
            x := float64(px)/width*(xmax-xmin) + xmin
            z := complex(x, y)
            // Punkt obrazu px, py reprezentuje wartosc zespolona z.
            img.Set(px, py, mandelbrot(z))
        }
    }
    png.Encode(os.Stdout, img) //UWAGA: ignorowanie bledow
}

func mandelbrot(z complex128) color.Color {
    const iterations = 200
    const contrast = 15

    var v complex128
    for n := uint8(0); n < iterations; n++ {
        v = v*v + z
        if cmplx.Abs(v) > 2 {
            nsmooth := float64(n) + 1 - math.Log(math.Log(math.Abs(float64(iterations))))/math.Log(2)
            c := uint32(nsmooth * 0xFFFFFFFF)
            r, g, b, a := uint8((c&0xFF000000)>>24), uint8((c&0x00FF0000)>>16), 
                    uint8((c&0x0000FF00)>>8), uint8(c&0x000000FF)
            return color.RGBA{r, g, b, a}
        }
    }
    return color.Black
}
