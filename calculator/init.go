package calculator

import (
	"github.com/yut-kt/speech2text/calculator/components"
)

func newContainers(frameNum int, components *components.Components) []*Container {
	words := components.GetWords()
	containers := make([]*Container, 0)
	for _, word := range words {
		table := make([][]*cell, frameNum)
		// 列情報の作成
		row := make([]*cell, 0)
		phonemeColumns := make([]string, 0)
		stateIndexColumns := make([]int, 0)
		for _, phoneme := range word.Phonemes {
			for stateIndex := 0; stateIndex < components.GetNumState(phoneme)-2; stateIndex++ {
				row = append(row, nil)
				phonemeColumns = append(phonemeColumns, phoneme)
				stateIndexColumns = append(stateIndexColumns, stateIndex)
			}
		}
		// 行情報(frame)の作成
		for i := 0; i < frameNum; i++ {
			table[i] = append(table[i], row...)
		}
		containers = append(containers, NewContainer(table, phonemeColumns, stateIndexColumns, word.Entry, word.Output))
	}
	return containers
}
