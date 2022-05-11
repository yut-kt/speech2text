package arpa

import (
	"bufio"
	"github.com/yut-kt/gostruct"
	"github.com/yut-kt/speech2text/util"
	"log"
)

type Arpa struct {
	unigramPs, unigramBackOffs map[string]float64
	bigramPs                   gostruct.Map2Dim[string, string, float64]
}

func NewArpa(arpaFile string) *Arpa {
	fp := util.OpenFile(arpaFile)
	defer util.CloseFile(fp)
	scanner := bufio.NewScanner(fp)
	unigramPs, unigramBackOffs, bigramPs := scanArpa(scanner)
	return &Arpa{
		unigramPs:       unigramPs,
		unigramBackOffs: unigramBackOffs,
		bigramPs:        bigramPs,
	}
}

func (a *Arpa) GetConnectP(w1, w2 string) float64 {
	//if w1 == ""

	if a.bigramPs.HasKey(w1, w2) {
		return a.bigramPs.Get(w1, w2)
	}

	if w2 == "</s>" { // EOSへの接続確率がないので定数を返す
		return -1000
	}

	p1, ok1 := a.unigramPs[w1]
	p2, ok2 := a.unigramBackOffs[w2]
	if !ok1 || !ok2 {
		log.Fatalf("ERROR: Not exitst [%v or %v]", w1, w2)
	}
	return p1 + p2/0.4342944819
}
