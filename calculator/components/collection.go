package components

import (
	"github.com/yut-kt/speech2text/calculator/components/arpa"
	"github.com/yut-kt/speech2text/calculator/components/htkdict"
	"github.com/yut-kt/speech2text/calculator/components/mmf"
	"math"
)

type Components struct {
	arpa *arpa.Arpa
	dict []*htkdict.Word
	mmf  *mmf.MMF
}

func NewComponents(arpaFile, dictFile, mmfFile string) *Components {
	return &Components{
		arpa: arpa.NewArpa(arpaFile),
		dict: htkdict.NewWords(dictFile),
		mmf:  mmf.NewMMF(mmfFile),
	}
}

func (c *Components) GetWords() []*htkdict.Word {
	return c.dict
}

func (c *Components) GetNumState(phoneme string) int {
	if phoneme == "silB" || phoneme == "silE" {
		phoneme = "sp"
	}
	return c.mmf.Phonemes[phoneme].NumState
}

func (c *Components) GetOutputP(feature []float64, phoneme string, stateIndex int) float64 {
	if phoneme == "silB" || phoneme == "silE" {
		phoneme = "sp"
	}
	stateInfo := c.mmf.Phonemes[phoneme].States[stateIndex]
	twoPi := 2.0 * math.Pi
	var o float64
	for i := 0; i < len(feature); i++ {
		m := stateInfo.Means[i]
		v := stateInfo.Variances[i]
		f := feature[i]
		o += math.Log(twoPi*v) + ((f - m) * (f - m) / v)
	}
	return -0.5 * o
}

func (c *Components) GetTransP(phoneme string, stateX, stateY int) float64 {
	if phoneme == "silB" || phoneme == "silE" {
		phoneme = "sp"
	}
	return c.mmf.Phonemes[phoneme].TransP[stateX][stateY]
}

func (c *Components) GetConnectionP(w1, w2 string) float64 {
	// TODO: 謎実装いる?
	//if w1 == "、+補助記号" || w2 == "、+補助記号" {
	//	return 0
	//}
	return c.arpa.GetConnectP(w1, w2) * 14
}
