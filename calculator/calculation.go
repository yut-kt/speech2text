package calculator

import (
	"fmt"
	"github.com/yut-kt/speech2text/calculator/components"
	"log"
	"math"
)

type Calculator struct {
	components *components.Components
	features   [][]float64
	containers []*Container
}

func NewCalculator(components *components.Components, features [][]float64) *Calculator {
	return &Calculator{
		components: components,
		features:   features,
		containers: newContainers(len(features), components),
	}
}

func (c *Calculator) Calculate() {
	const beamW = 200

	// 1フレーム目だけ特別
	p := c.calcOutputPBOS()
	c.calcSelfTransP(0, p-beamW)
	c.calcNextTransP(0, p-beamW)
	c.calcConnectP(0, p-beamW)

	for i := 1; i < len(c.features)-1; i++ {
		p := c.calcOutputP(i)
		limit := p - beamW
		c.calcSelfTransP(i, limit)
		c.calcNextTransP(i, limit)
		c.calcConnectP(i, limit)
	}

	result := make([]*cell, 0)
	for _, container := range c.containers {
		if container.entry == "</s>" {
			c := container.table[len(container.table)-1][len(container.table[0])-1]
			result = append(result, c)
			for c = c.backcell; c != nil; c = c.backcell {
				if result[len(result)-1].origin != c.origin {
					result = append(result, c)
				}
			}
			break
		}
	}
	for i := len(result) - 1; i >= 0; i-- {
		fmt.Print(result[i].origin.output)
		fmt.Print(" ")
	}
	fmt.Println()
}

func (c *Calculator) calcOutputPBOS() float64 {
	for _, container := range c.containers {
		if container.entry == "<s>" {
			outputP := c.components.GetOutputP(c.features[0], container.phonemeColumns[0], container.stateIndexColumns[0])
			container.table[0][0] = newCell(container, outputP, nil)
			return outputP
		}
	}
	log.Fatal("ERROR: Not found BOS.")
	return 1
}

func (c *Calculator) calcOutputP(frameIndex int) float64 {
	max := -math.MaxFloat64
	for _, container := range c.containers {
		for i := 0; i < len(container.table[0]); i++ {
			if container.table[frameIndex][i] != nil {
				container.table[frameIndex][i].p += c.components.GetOutputP(c.features[frameIndex], container.phonemeColumns[i], container.stateIndexColumns[i])
				if container.table[frameIndex][i].p > max {
					max = container.table[frameIndex][i].p
				}
			}
		}
	}
	return max
}

func (c *Calculator) calcSelfTransP(frameIndex int, limit float64) {
	for _, container := range c.containers {
		for i := 0; i < len(container.table[0]); i++ {
			if cell := container.table[frameIndex][i]; cell != nil && cell.p > limit {
				phoneme := container.phonemeColumns[i]
				stateIndex := container.stateIndexColumns[i]
				transP := c.components.GetTransP(phoneme, stateIndex, stateIndex+1)
				nextP := cell.p + transP
				if nextCell := container.table[frameIndex+1][i]; nextCell == nil || nextP > nextCell.p {
					container.table[frameIndex+1][i] = newCell(container, nextP, cell)
				}
			}
		}
	}
}

func (c *Calculator) calcNextTransP(frameIndex int, limit float64) {
	for _, container := range c.containers {
		for i := 0; i < len(container.table[0])-1; i++ {
			if cell := container.table[frameIndex][i]; cell != nil && cell.p > limit {
				phoneme := container.phonemeColumns[i]
				stateIndex := container.stateIndexColumns[i]
				transP := c.components.GetTransP(phoneme, stateIndex+1, stateIndex+1)
				nextP := cell.p + transP
				if nextCell := container.table[frameIndex+1][i+1]; nextCell == nil || nextP > nextCell.p {
					container.table[frameIndex+1][i+1] = newCell(container, nextP, cell)
				}
			}
		}
	}
}

func (c *Calculator) calcConnectP(frameIndex int, limit float64) {
	for _, container := range c.containers {
		if container.entry == "</s>" { // EOSからは出られない
			continue
		}

		lastColumnIndex := len(container.table[0]) - 1
		if lastColumnCell := container.table[frameIndex][lastColumnIndex]; lastColumnCell != nil && lastColumnCell.p > limit {
			phoneme := container.phonemeColumns[lastColumnIndex]
			stateIndex := container.stateIndexColumns[lastColumnIndex]
			transP := c.components.GetTransP(phoneme, stateIndex+1, stateIndex+1)
			nextP := lastColumnCell.p + transP
			for _, nextContainer := range c.containers {
				if nextContainer.entry == "<s>" {
					continue
				}
				connectP := c.components.GetConnectionP(container.entry, nextContainer.entry) + nextP
				if connectCell := nextContainer.table[frameIndex+1][0]; connectCell == nil || connectP > connectCell.p {
					nextContainer.table[frameIndex+1][0] = newCell(nextContainer, connectP, lastColumnCell)
				}
			}
		}
	}
}
