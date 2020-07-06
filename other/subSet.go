package other

import (
	"fmt"
)

var sets [4]string = [4]string{"a", "b", "c", "d"}
var size int = 4
var isVisited [4]bool = [4]bool{false}

// getSubSets ...
func getSubSets(depth int) {
	if depth == size {
		for i := 0; i < size; i++ {
			if isVisited[i] {
				fmt.Print(sets[i])
			}
		}
		fmt.Println()

	} else {
		isVisited[depth] = true
		getSubSets(depth + 1)
		isVisited[depth] = false
		getSubSets(depth + 1)
	}
}

