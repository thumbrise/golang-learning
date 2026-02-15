package tower_of_hanoi

import "github.com/emirpasic/gods/stacks/arraystack"

type Tower struct {
	Name  string
	Disks *arraystack.Stack
}

func TowerOfHanoi(origin, destination, auxiliary *Tower) {
	value, _ := origin.Disks.Pop()
	if origin.Disks.Size() <= 1 {
		destination.Disks.Push(value)
		return
	}

	TowerOfHanoi(origin, auxiliary, destination)
	destination.Disks.Push(value)
	TowerOfHanoi(destination, origin, auxiliary)
}
