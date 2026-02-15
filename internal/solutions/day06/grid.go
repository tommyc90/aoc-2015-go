package day06

type grid [][]bool

func newGrid(rows, cols int) grid {
	g := make(grid, rows)
	for i := range rows {
		g[i] = make([]bool, cols)
	}
	return g
}

func (g grid) applyInstruction(ins instruction) {
	for ri := ins.fromRow; ri <= ins.toRow; ri++ {
		row := g[ri]
		for ci := ins.fromCol; ci <= ins.toCol; ci++ {
			switch ins.mode {
			case on:
				row[ci] = true
			case off:
				row[ci] = false
			case toggle:
				row[ci] = !row[ci]
			}
		}
	}
}

func (g grid) countLights() int {
	c := 0
	for _, row := range g {
		for _, v := range row {
			if v {
				c++
			}
		}
	}
	return c
}
