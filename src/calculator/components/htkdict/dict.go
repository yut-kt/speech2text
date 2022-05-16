package htkdict

import (
	"bufio"

	"github.com/yut-kt/speech2text/src/util"
)

type Word struct {
	Entry    string
	Output   string
	Phonemes []string
}

func NewWords(dictFile string) []*Word {
	fp := util.OpenFile(dictFile)
	defer util.CloseFile(fp)
	scanner := bufio.NewScanner(fp)
	return ScanWords(scanner)
}
