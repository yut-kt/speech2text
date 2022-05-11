package arpa

import (
	"bufio"
	"github.com/yut-kt/gostruct"
	"github.com/yut-kt/speech2text/util"
	"log"
	"strings"
)

func scanArpa(scanner *bufio.Scanner) (pMap map[string]float64, bMap map[string]float64, map2D gostruct.Map2Dim[string, string, float64]) {
	for scanner.Scan() {
		switch scanner.Text() {
		case "\\1-grams:":
			pMap, bMap = scanUnigram(scanner)
		case "\\2-grams:":
			map2D = scanBigram(scanner)
			return
		case "\\3-grams:":
			//ngram.ReadTrigrams(scanner)
		}
	}
	return
}

func scanUnigram(scanner *bufio.Scanner) (pMap map[string]float64, bMap map[string]float64) {
	pMap, bMap = make(map[string]float64, 0), make(map[string]float64, 0)
	const silE = "</s>"
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		lineTexts := strings.Split(scanner.Text(), "\t")
		p := util.StringToFloat64(lineTexts[0])
		word := lineTexts[1]

		if word == silE { // EOSのみバックオフ係数が存在しない
			pMap[silE] = p
		} else {
			pMap[word] = p
			bMap[word] = util.StringToFloat64(lineTexts[2])
		}
	}
	return
}

func scanBigram(scanner *bufio.Scanner) (m2D gostruct.Map2Dim[string, string, float64]) {
	m2D = make(gostruct.Map2Dim[string, string, float64], 0)
	for scanner.Scan() {
		if scanner.Text() == "" {
			return
		}

		lineTexts := strings.Split(scanner.Text(), "\t")
		p := util.StringToFloat64(lineTexts[0])
		words := strings.Split(lineTexts[1], " ")
		if len(words) != 2 {
			log.Fatalf("ERROR: not correct words len. len(words) = [%v]", len(words))
		}
		m2D.Set(words[0], words[1], p)
	}
	return
}
