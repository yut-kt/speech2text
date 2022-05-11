package feature

import (
	"github.com/yut-kt/gowave"
	"github.com/yut-kt/speech2text/util"
	"log"
)

func ReadWave(waveFile string) ([]float64, int) {
	fp := util.OpenFile(waveFile)
	defer util.CloseFile(fp)

	wave, err := gowave.New(fp)
	if err != nil {
		log.Fatal(err)
	}

	samples, err := wave.ReadSamples()
	if err != nil {
		log.Fatal(err)
	}

	var f64s []float64
	switch slice := samples.(type) {
	case []uint8:
	case []int16:
		f64s = make([]float64, len(slice))
		for i, v := range slice {
			f64s[i] = float64(v)
		}
	}
	return f64s, int(wave.GetSampleRate())
}
