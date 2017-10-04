package main

import (
    "fmt"
    "log"
    "net/http"
    "image/color"
    "math"
)

const (
    width, height = 600, 320
    cells = 100
    xyrange = 30.0
    xyscale = width / 2 / xyrange
    zscale = height * 0.4
    angle = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "image/svg+xml")
        surface(w)
    })
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}


func surface(w http.ResponseWriter) {
    fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
        "style='stroke: grey; fill: white; stroke-width: 0.7' "+
        "width='%d' height='%d'>", width, height)
    for i := 0; i < cells; i++ {
        for j := 0; j < cells; j++ {
            var z1, z2, z3, z4 float64
            var ax, ay, bx, by, cx, cy, dx, dy float64
            style := "white"
            ax, ay, z1 = corner(i+1, j)
            if math.IsNaN(z1) {
                continue
            }
            bx, by, z2 = corner(i, j)
            if math.IsNaN(z2) {
                continue
            }
            cx, cy, z3 = corner(i, j+1)
            if math.IsNaN(z3) {
                continue
            }
            dx, dy, z4 = corner(i+1, j+1)
            if math.IsNaN(z4) {
                continue
            }
            z := (z1+z2+z3+z4) / 4
            if z > 0 {
                style = "red"
            }
            if z < 0 {
                style = "blue"
            }

            fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill: %s' />\n",
                ax, ay, bx, by, cx, cy, dx, dy, style)
        }
    }
    fmt.Fprintf(w, "</svg>\n")
}

func corner(i, j int) (float64, float64, float64) {
    // Znajdowanie x,y, w rogu komorki i,j
    x := xyrange * (float64(i)/cells - 0.5)
    y := xyrange * (float64(j)/cells - 0.5)

    // Obliczanie wysokosci z powierzchni
    z := f(x, y)

    // Rzutowanie x,y,z izometrycznie na plotno 2D SVG (sx, sy)
    sx := width/2 + (x-y)*cos30*xyscale
    sy := height/2 + (x+y)*sin30*xyscale - z*zscale
    return sx, sy, z
}

func f(x, y float64) float64 {
    r := math.Hypot(x, y) // odleglosc od punktu 0,0
    return math.Sin(r) / r
}

// Wygeneruj tube 
func tube(x, y float64) float64 {
    return 1 / (15*(x*x + y*y)) 
}

func bumps(x, y float64) float64 {
    return math.Sin(5*x)*math.Cos(5*y)/5
}
