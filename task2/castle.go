package task2

type coords struct {
	x    int
	y    int
	step int
}

func (c *coords) checkBoundaries(m, n int) bool {
	return c.x >= 0 && c.y >= 0 && c.x < m && c.y < n
}

func (c *coords) compare(cc *coords) bool {
	return c.x == cc.x && c.y == cc.y
}

// Room is technically a value in a Castle matrix, val defines whether a room is blocked or open, the values are 0 and 1 consequently.
// Also, 0 value for the val field means that the room has been visited.
type Room struct {
	val int
}

// SetValue creates new room
func (r *Room) SetValue(val int) {
	r.val = val
}

func (r *Room) setBlocked() {
	r.val = 0
}

func (r *Room) isOpen() bool {
	return r.val == 1
}

// Castle is a matrix of M x N dimensions with 0s as blocked rooms and 1s as passable rooms.
type Castle struct {
	m     int
	n     int
	Rooms [][]Room
}

// NewCastle creates a new castle with empty rooms
func NewCastle(m, n int) *Castle {
	castle := Castle{m, n, make([][]Room, m)}
	for i := 0; i < castle.m; i++ {
		castle.Rooms[i] = make([]Room, n)
	}
	return &castle
}

var directions = []coords{{0, 1, 1}, {1, 0, 1}, {0, -1, 1}, {-1, 0, 1}}

// FindPathCount is a function that calculates number of possible paths from bottom left to top right corner
func (c *Castle) FindPathCount() int {
	allPathsNum := 0
	exit := coords{0, c.n - 1, 0}
	queue := make([]coords, 0)

	// Bottom left room is first to explore
	queue = append(queue, coords{c.m - 1, 0, 1})

	for len(queue) > 0 {
		// Remove first coords from the queue and assign them to a temp variable that implies current position
		currentPos := coords{queue[0].x, queue[0].y, queue[0].step}
		queue = queue[1:]

		c.Rooms[currentPos.x][currentPos.y].setBlocked()

		// Check if current position is on exit
		if currentPos.compare(&exit) && currentPos.step <= (c.m+c.n-1) {
			allPathsNum++
		}

		for _, d := range directions {
			nextPos := coords{currentPos.x + d.x, currentPos.y + d.y, currentPos.step + 1}

			if nextPos.checkBoundaries(c.m, c.n) && c.Rooms[nextPos.x][nextPos.y].isOpen() {
				queue = append(queue, nextPos)
			}
		}
	}

	return allPathsNum
}
