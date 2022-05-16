package mmf

import (
	"bufio"
	"log"
	"strconv"
	"strings"

	"github.com/yut-kt/speech2text/src/util"
)

type Phoneme struct {
	NumState int
	States   []*State
	TransP   [][]float64
}

func NewPhoneme(scanner *bufio.Scanner) *Phoneme {
	phoneme := &Phoneme{}

	// BEGIN HMM
	scanner.Scan()
	checkBegin(scanner.Text())
	// State
	scanner.Scan()
	phoneme.NumState = checkNumState(scanner.Text())
	phoneme.States = make([]*State, phoneme.NumState-2)
	for index := 0; index < phoneme.NumState-2; index++ {
		scanner.Scan()
		checkState(scanner.Text(), index)
		phoneme.States[index] = NewState(scanner)
	}
	// TransP
	scanner.Scan()
	numTransP := checkTransP(scanner.Text())
	phoneme.TransP = readTransP(scanner, numTransP)

	scanner.Scan()
	checkEnd(scanner.Text())

	return phoneme
}

func checkBegin(s string) {
	if s != "<BEGINHMM>" {
		log.Fatalf("ERROR: Not found <BEGINHMM>. found %v.", s)
	}
}

func checkNumState(s string) int {
	ss := strings.Split(s, " ")
	if ss[0] != "<NUMSTATES>" {
		log.Fatalf("ERROR: Not found <NUMSTATES>. found %v.", ss[0])
	}
	numState, err := strconv.Atoi(ss[1])
	if err != nil {
		log.Fatal(err)
	}
	return numState
}

func checkState(s string, i int) {
	ss := strings.Split(s, " ")
	if ss[0] != "<STATE>" {
		log.Fatalf("ERROR: Not found <STATE>. found %v", ss[0])
	}
	stateNum, err := strconv.Atoi(ss[1])
	if err != nil {
		log.Fatal(err)
	}
	if i != stateNum-2 {
		log.Fatalf("ERROR: difference state. i[%v] != stateNumber[%v]", i, stateNum-2)
	}
}

func checkTransP(s string) int {
	ss := strings.Split(s, " ")
	if ss[0] != "<TRANSP>" {
		log.Fatalf("ERROR: Not found <TRANSP>. found %v.", ss[0])
	}
	numTransP, err := strconv.Atoi(ss[1])
	if err != nil {
		log.Fatal(err)
	}
	return numTransP
}

func readTransP(scanner *bufio.Scanner, numTransP int) [][]float64 {
	transP := make([][]float64, numTransP)
	for i := 0; i < numTransP; i++ {
		transP[i] = make([]float64, numTransP)
	}
	for i := range transP {
		scanner.Scan()
		transP[i] = util.StringLineToFloat64s(scanner.Text(), numTransP)
	}
	return transP
}

func checkEnd(s string) {
	if s != "<ENDHMM>" {
		log.Fatalf("ERROR: Not found <ENDHMM>. found %v.", s)
	}
}
