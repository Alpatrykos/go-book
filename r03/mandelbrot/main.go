package main

import (
    "image"
    "image/color"
    "image/png"
    "math/cmplx"
    "os"
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
    const contrast = 10

    var v complex128
    for n := uint8(0); n < iterations; n++ {
        v = v*v + z
        if cmplx.Abs(v) > 2 {
            r, b, g := uint8(0), uint8(0), uint8(0)
            a := uint8(255)
            sum := 3*255
            sum -= int(n)*contrast
            switch sum/255 {
            case 3:
                r,b,g = 255, 255, 255
            case 2:
                r,b = 255, 255
                g = uint8(sum % 255)
            case 1:
                r = 255
                g = 0
                b = uint8(sum % 255)
            default:
                r,b,g = 0, 0, 0
            }
            return color.RGBA{r, b, g, a}
        }
    }
    return color.Black
}
