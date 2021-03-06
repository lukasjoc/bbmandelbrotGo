package bbmandelbrot

import (
	"flag"
	"fmt"
	"image/png"
	"os"
	"runtime"
	"testing"
	"time"
)

var (
	cscheme string
	fname   string
	width   uint64
	height  uint64
	cx1     uint64
	cx2     uint64
	cy1     uint64
	cy2     uint64
	csr     int
	csg     int
	csb     int
)

func Test_Mandelbrot(t *testing.T) {
	start := time.Now()
	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.StringVar(&fname, "f", "mandelbrot.png", "destination filename")
	flag.Uint64Var(&width, "w", 2560, "fractal width")
	flag.Uint64Var(&height, "h", 2560, "fractal height")
	flag.Uint64Var(&cx1, "cx1", 0, "crop width start")
	flag.Uint64Var(&cx2, "cx2", 0, "crop width end")
	flag.Uint64Var(&cy1, "cy1", 0, "crop height start")
	flag.Uint64Var(&cy2, "cy2", 0, "crop height end")
	flag.IntVar(&csr, "r", 2, "color scheme (red)")
	flag.IntVar(&csg, "g", 3, "color scheme (green)")
	flag.IntVar(&csb, "b", 1, "color scheme (blue)")
	flag.Parse()

	if cx2 == 0 {
		cx2 = width
	}
	if cy2 == 0 {
		cy2 = height
	}

	img, retstr := Mandelbrot(width, height, cx1, cx2, cy1, cy2, csr, csg, csb)
	fmt.Println(retstr)

	file, err := os.Create(fname)
	defer file.Close()

	if err != nil || file == nil {
		file, err = os.Open(fname)
		defer file.Close()
		if err != nil {
			t.Error(fmt.Sprintf("Error opening file: %s\n", err))
		}
	}

	err = png.Encode(file, img)
	if err != nil {
		t.Error(fmt.Sprintf("Error encoding image: %s\n", err))
	}
	fmt.Printf("\033[2Jimage saved to %v after %v\n", fname, time.Since(start))
}
