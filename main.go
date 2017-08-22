package main

//import "fmt"
import (
	. "./sub"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

func main() {
	//runtime.GOMAXPROCS(runtime.NumCPU())
	// log.Println(runtime.NumCPU())
	var argv []string = os.Args
	var argc int = len(os.Args)

	var W int = 16

	if argc == 2 {
		W, _ = strconv.Atoi(argv[1])
	}

	log.Println(runtime.GOMAXPROCS(0))
	// log.Println(runtime.NumGoroutine())
	matA := make([]float64, W*W)
	matB := make([]float64, W*W)
	matC := make([]float64, W*W)
	matCp := make([]float64, W*W)
	rand.Seed(time.Now().UnixNano())

	ts := NewTimeStamp(10)
	//ts.Stamp();

	for i := 0; i < W*W; i++ {
		matA[i] = rand.Float64()
		matB[i] = rand.Float64()
		//fmt.Printf("%f %f\n",matA[i],matB[i])
	}

	fmt.Printf("Matrix Multiplication %dx%d\n", W, W)
	ts.Stamp()

	for w := 0; w < W; w++ {
		for h := 0; h < W; h++ {
			var s float64 = 0
			for k := 0; k < W; k++ {
				s += matA[k+h*W] * matB[w+k*W]
			}
			matC[w+h*W] = s
		}
	}

	ts.Stamp()
	var wg sync.WaitGroup
	for w := 0; w < W; w++ {
		for h := 0; h < W; h++ {

			wg.Add(1)
			go func(w int, h int) {
				defer wg.Done()
				var s float64 = 0
				for k := 0; k < W; k++ {
					s += matA[k+h*W] * matB[w+k*W]
				}
				matCp[w+h*W] = s
			}(w, h)


		}
	}
	wg.Wait()
	ts.Stamp()
	// ts.Print()
	// ts.Print_str([]string{"Sequential", "Parallel"})
	// ts.Print_oneLine()
	ts.Print_str_oneLine([]string{"Sequential", "Parallel"})

}
