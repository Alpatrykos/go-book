package main

import (
    "image"
    "image/color"
    "image/png"
    "math/cmplx"
    "os"
    //"math"
)

func main() {
    const (
        xmin, ymin, xmax, ymax = -2, -2, +2, +2
        width, height = 4096, 4096
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
    rimg := supersample2(img)
    png.Encode(os.Stdout, rimg) //UWAGA: ignorowanie bledow
}

func mandelbrot(z complex128) color.Color {
    const iterations = 200
    const contrast = 15

    var v complex128
    for n := uint8(0); n < iterations; n++ {
        v = v*v + z
        if cmplx.Abs(v) > 2 {
            //nsmooth := float64(n) + 1 - math.Log(math.Log(math.Abs(float64(iterations))))/math.Log(2)
            //nsmooth /= iterations
            //c := uint32(nsmooth * 0x00FFFFFF)
            //r, g, b := uint8((c&0x00FF0000)>>16), uint8((c&0x0000FF00)>>8), 
                    //uint8((c&0x000000FF))
            return color.RGBA{contrast * (n+2), contrast * (n+1), contrast * n, 255}
        }
    }
    return color.Black
}

// does not work as intended; the image actually becomes more aliased in a result of supersampling
// method implemented in the fuction.
func supersample(img image.Image) image.Image {
    bounds := img.Bounds()
    ssimg := image.NewRGBA(image.Rect(0, 0, bounds.Max.X/2, bounds.Max.Y/2))
    for py := 0; py < ssimg.Bounds().Max.Y; py++ {
        for px := 0; px < ssimg.Bounds().Max.X; px++ {
           r1, g1, b1, _ := img.At(2*px-1, 2*py-1).RGBA()
           r2, g2, b2, _ := img.At(2*px, 2*py-1).RGBA()
           r3, g3, b3, _ := img.At(2*px-1, 2*py).RGBA()
           r4, g4, b4, _ := img.At(2*px, 2*py).RGBA()
           avg_red := float64(r1 + r2 + r3 + r4)
           avg_green := float64(g1 + g2 + g3 + g4)
           avg_blue := float64(b1 + b2 + b3 + b4)
           c := color.RGBA{uint8(avg_red/4), uint8(avg_green/4), uint8(avg_blue/4), 255}
           ssimg.Set(px, py, c)
       }
   }
   return ssimg
}

func supersample2(img image.Image) image.Image {
    bounds := img.Bounds()
    ssimg := image.NewRGBA(image.Rect(0, 0, bounds.Max.X-1, bounds.Max.Y-1))
    for py := 0; py < ssimg.Bounds().Max.Y; py++ {
        for px := 0; px < ssimg.Bounds().Max.X; px++ {
            r1, g1, b1, _ := img.At(px, py).RGBA()
            r2, g2, b2, _ := img.At(px+1, py).RGBA()
            r3, g3, b3, _ := img.At(px, py+1).RGBA()
            r4, g4, b4, _ := img.At(px+1, py+1).RGBA()
            avg_r := float64(r1 + r2 + r3 + r4)
            avg_g := float64(g1 + g2 + g3 +g4)
            avg_b := float64(b1 + b2 + b3 + b4)
            c := color.RGBA{uint8(avg_r/4), uint8(avg_g/4), uint8(avg_b/4), 255}
            ssimg.Set(px, py, c)
        }
    }
    return ssimg
}
