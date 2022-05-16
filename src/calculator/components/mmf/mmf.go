package mmf

import (
	"bufio"
	"regexp"
	"strings"

	"github.com/yut-kt/speech2text/src/util"
)

type MMF struct {
	Phonemes map[string]*Phoneme
}

func NewMMF(mmfFile string) (mmf *MMF) {
	mmf = &MMF{map[string]*Phoneme{}}

	fp := util.OpenFile(mmfFile)
	defer util.CloseFile(fp)
	scanner := bufio.NewScanner(fp)

	regex := regexp.MustCompile("\"(.*)\"")
	for scanner.Scan() {
		line := scanner.Text()
		if line[:2] == "~h" {
			phoneme := strings.ReplaceAll(regex.FindString(line), "\"", "")
			mmf.Phonemes[phoneme] = NewPhoneme(scanner)
		}
	}
	return
}
