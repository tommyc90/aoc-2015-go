package day06

type grid2 [][]int

func newGrid2(x, y int) grid2 {
	g := make(grid2, x)
	for i := range x {
		g[i] = make([]int, y)
	}
	return g
}

func (g grid2) applyInstruction(ins instruction) {
	for ri := ins.fromRow; ri <= ins.toRow; ri++ {
		row := g[ri]
		for ci := ins.fromCol; ci <= ins.toCol; ci++ {
			switch ins.mode {
			case on:
				row[ci]++
			case off:
				if row[ci] > 0 {
					row[ci]--
				}
			case toggle:
				row[ci] += 2
			}
		}
	}
}

func (g grid2) brightness() int {
	b := 0
	for _, row := range g {
		for _, v := range row {
			b += v
		}
	}
	return b
}
