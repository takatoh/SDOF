package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/takatoh/sdof/directintegration"
	"github.com/takatoh/seismicwave"
)

func main() {
	methods := []string{
		"wilson-theta",
		"average",
		"linear",
		"nigam",
	}

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
	opt_method := flag.String("method", "wilson-theta", "Specify integration method. default to wilson-theta.")
	flag.Parse()

	if !(contains(methods, *opt_method)) {
		fmt.Fprintf(os.Stderr, "Error: Unsupported integration method: %s\n", *opt_method)
		fmt.Fprintf(os.Stderr, "Available methods are: %s\n", strings.Join(methods, ", "))
		os.Exit(1)
	}
	filename := flag.Arg(0)

	var waves []*seismicwave.Wave
	var err error
	waves, err = seismicwave.LoadCSV(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	ndiv := 10
	wave := waves[0]
	data := interpolate(wave.Data, ndiv)
	n := len(data)
	dt := wave.DT() / float64(ndiv)
	w := *opt_omega
	h := *opt_h

	var acc []float64
	if *opt_method == "wilson-theta" {
		acc, _, _ = directintegration.WilsonTheta(h, w, dt, n, data)
	} else if *opt_method == "average" {
		acc, _, _ = directintegration.AverageAcc(h, w, dt, n, data)
	} else if *opt_method == "linear" {
		acc, _, _ = directintegration.LinearAcc(h, w, dt, n, data)
	} else if *opt_method == "nigam" {
		acc, _, _ = directintegration.Nigam(h, w, dt, n, data)
	}
	fmt.Println("Time,Acc")
	t := 0.0
	dt = wave.DT()
	for i := 0; i < len(acc); i += ndiv {
		fmt.Printf("%f,%f\n", t, acc[i])
		t = t + dt
	}
}

func interpolate(zin []float64, ndiv int) []float64 {
	var zinc float64
	nin := len(zin)
	ndivf := float64(ndiv)
	z := make([]float64, 0)
	k := 0
	z = append(z, 0.0)
	for i := 0; i < nin; i++ {
		if i == 0 {
			zinc = zin[i] / ndivf
		} else {
			zinc = (zin[i] - zin[i-1]) / ndivf
		}
		for j := 0; j < ndiv; j++ {
			z = append(z, z[k]+zinc)
			k++
		}
	}
	return z
}

func contains(elem []string, v string) bool {
	for _, s := range elem {
		if v == s {
			return true
		}
	}
	return false
}
