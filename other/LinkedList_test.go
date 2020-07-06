package other

import (
	"fmt"
	"testing"
)

// TestLinkedList ...
func TestLinkedList(t *testing.T)  {
	//insertNode(1)
	//insertNode(2)
	//insertNode(3)
	//insertNode(4)
	//insertNode(5)
	for i := 10; i > 0; i-- {
		insertNode(i)
	}
	display()
	fmt.Println(findNode(3))
	remove(3)
	fmt.Println(findNode(3))
	display()
}
