package main

import (
	"fmt"
	"os"

	c "github.com/Highzest/interview-task/task2"
)

func main() {
	var m int
	var n int

	_, err := fmt.Scan(&m, &n)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error occured when parsing m and n from stdin: %v\n", err)
		os.Exit(1)
	}

	if m < 0 || m > 1000 || n < 0 || n > 1000 {
		fmt.Fprintf(os.Stderr, "Error: m or n can be between 0 and 1000. You entered %d and %d\n", m, n)
		os.Exit(1)
	}

	castle := c.NewCastle(m, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			var val int
			if b, err := fmt.Scan(&val); b != 1 {
				fmt.Fprintf(os.Stderr, "Error occured when parsing the matrix from stdin: %v\n", err)
				os.Exit(1)
			}

			if val < 0 && val > 1 {
				fmt.Fprintf(os.Stderr, "Error: only 1 or 0 are allowed. You entered %d\n", val)
				os.Exit(1)
			}

			(castle.Rooms[i][j]).SetValue(val)
		}
	}

	for i := 0; i < m; i++ {
		fmt.Println(castle.Rooms[i])
	}

	paths := castle.FindPathCount()
	if paths != 0 {
		fmt.Println(paths)
	} else {
		fmt.Println("impossible")
	}
}
