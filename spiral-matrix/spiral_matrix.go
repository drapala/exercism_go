package spiralmatrix

type matrix struct {
	dim       int
	grid      [][]int
	x         int    // x index
	y         int    // y index
	direction string // starting direction
}

// Initialize the matrix with all zeros
func (m *matrix) initZeros() {
	grid := make([][]int, m.dim)
	// Create rows
	for i := 0; i < m.dim; i++ {
		grid[i] = make([]int, m.dim)
		m.grid = append(m.grid, grid[i])
	}
}

// Move logic
func (m *matrix) move() {
	switch m.direction {
	case "right":
		m.moveRight()
	case "left":
		m.moveLeft()
	case "up":
		m.moveUp()
	case "down":
		m.moveDown()
	}
	return
}

func (m *matrix) moveDown() {
	// If less than size and next one is zero
	if m.y < m.dim-1 && m.grid[m.y+1][m.x] == 0 {
		m.y++
		return // Move successful
	}
	// At the edge, direction needs to change
	m.direction = "left"
	m.x--
	return // At the edge
}

func (m *matrix) moveUp() {
	// If less than size and next one is zero
	if m.y > 0 && m.grid[m.y-1][m.x] == 0 {
		m.y--
		return // Move successful
	}
	// At the edge, direction needs to change
	m.direction = "right"
	m.x++
	return // At the edge
}

func (m *matrix) moveLeft() {
	// If less than size and next one is zero
	if m.x > 0 && m.grid[m.y][m.x-1] == 0 {
		m.x--
		return // Move successful
	}
	// At the edge, direction needs to change
	m.direction = "up"
	m.y--
	return // At the edge
}

func (m *matrix) moveRight() {
	// If less than size and next one is zero
	if m.x < m.dim-1 && m.grid[m.y][m.x+1] == 0 {
		m.x++
		return // Move successful
	}
	// At the edge, direction needs to change
	m.direction = "down"
	m.y++
	return // At the edge
}

func SpiralMatrix(size int) [][]int {
	if size == 0 { // Error case
		return [][]int{}
	}
	myGrid := matrix{dim: size, x: 0, y: 0, direction: "right"} // Starting direction is right
	myGrid.initZeros()
	num := 1 // Will go up to size * size
	for num <= size*size {
		myGrid.grid[myGrid.y][myGrid.x] = num
		myGrid.move()
		num++
	}
	return myGrid.grid
}
