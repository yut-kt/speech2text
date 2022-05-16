package mmf

import (
	"bufio"
	"log"
	"strconv"
	"strings"

	"github.com/yut-kt/speech2text/src/util"
)

type State struct {
	Means     []float64
	Variances []float64
	GConst    float64
}

func NewState(scanner *bufio.Scanner) (state *State) {
	state = &State{}
	for scanner.Scan() {
		ss := strings.Split(scanner.Text(), " ")
		num, err := strconv.ParseFloat(ss[1], 64)
		if err != nil {
			log.Fatal(err)
		}

		switch ss[0] {
		case "<MEAN>":
			scanner.Scan()
			state.Means = util.StringLineToFloat64s(scanner.Text(), int(num))
		case "<VARIANCE>":
			scanner.Scan()
			state.Variances = util.StringLineToFloat64s(scanner.Text(), int(num))
		case "<GCONST>":
			state.GConst = num
			return // gconstまで読んだらstate終了
		default:
			log.Fatalf("ERROR: Unknown attr [%v].", ss[0])
		}
	}
	log.Fatal("ERROR: Not read <GCONST>.")
	return
}
