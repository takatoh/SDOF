package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/takatoh/sdof/directintegration"
	"github.com/takatoh/seismicwave"
)

func main() {
	flag.Parse()

	filename := flag.Arg(0)

	var waves []*seismicwave.Wave
	var err error
	waves, err = seismicwave.LoadCSV(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	wave := waves[0]
	n := wave.NData()
	dt := wave.DT()
	data := wave.Data
	w := 1.0
	h := 0.05

	acc, _, _ := directintegration.WilsonTheta(w, h, dt, n, data)
	t := 0.0
	for i := 0; i < len(acc); i++ {
		fmt.Printf("%f,%f\n", t, acc[i])
		t = t + dt
	}
}
