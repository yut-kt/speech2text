package main

import (
	"bufio"
	"fmt"
	"github.com/yut-kt/speech2text/calculator"
	"github.com/yut-kt/speech2text/calculator/components"
	"github.com/yut-kt/speech2text/util"
)

func main() {
	collection := components.NewComponents(
		"storage/defs/bccwj.60k.tri.arpa",
		"storage/defs/bccwj.60k.htkdic.small",
		"storage/defs/hmmdefs",
	)

	// 特徴量によってEOSまで辿りつかない場合がある．
	//samples, sampleRate := feature.ReadWave("storage/wav/nitech_jp_atr503_m001_a01.wav")
	//features2 := feature.GetMFCC(samples, sampleRate)
	//fmt.Println(len(features2), len(features2[0]))

	var features [][]float64
	fp := util.OpenFile("storage/feature")
	defer util.CloseFile(fp)
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		features = append(features, util.StringLineToFloat64s(scanner.Text(), 39))
	}
	fmt.Println(len(features), len(features[0]))

	calculator.NewCalculator(collection, features).Calculate()
}
