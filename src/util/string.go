package util

import (
	"log"
	"strconv"
	"strings"
)

func StringLineToFloat64s(s string, n int) (result []float64) {
	ss := strings.Split(strings.TrimSpace(s), " ")
	if len(ss) != n {
		log.Fatalf("ERROR: Size is different. s[%v] != n[%v]", len(ss), n)
	}

	result = make([]float64, n)
	for i := range ss {
		f64, err := strconv.ParseFloat(ss[i], 64)
		if err != nil {
			log.Fatal(err)
		}
		result[i] = f64
	}
	return
}

func StringToFloat64(s string) float64 {
	f64, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatal(err)
	}
	return f64
}
