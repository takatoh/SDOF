package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/takatoh/sdof/directintegration"
	"github.com/takatoh/seismicwave"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr,
			`Usage:
  %s [options] <wavefile.csv>

Options:
`, os.Args[0])
		flag.PrintDefaults()
	}
	opt_omega := flag.Float64("omega", 1.0, "Specify sircular frequency.")
	opt_h := flag.Float64("h", 0.05, "Specify attenuation constant.")
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
	w := *opt_omega
	h := *opt_h

	acc, _, _ := directintegration.WilsonTheta(h, w, dt, n, data)
	fmt.Println("Time,Acc")
	t := 0.0
	for i := 0; i < len(acc); i++ {
		fmt.Printf("%f,%f\n", t, acc[i])
		t = t + dt
	}
}
