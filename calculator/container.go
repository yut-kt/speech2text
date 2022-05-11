package calculator

type Container struct {
	table             [][]*cell
	phonemeColumns    []string
	stateIndexColumns []int
	entry             string
	output            string
}

func NewContainer(table [][]*cell, phonemeColumns []string, stateIndexColumns []int, entry, output string) *Container {
	return &Container{table: table, phonemeColumns: phonemeColumns, stateIndexColumns: stateIndexColumns, entry: entry, output: output}
}

type cell struct {
	origin   *Container
	p        float64
	backcell *cell
}

func newCell(origin *Container, p float64, backcell *cell) *cell {
	return &cell{origin: origin, p: p, backcell: backcell}
}
