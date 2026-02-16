package tower_of_hanoi

import (
	"github.com/emirpasic/gods/stacks/arraystack"
)

type Tower struct {
	Name  string
	Disks *arraystack.Stack
}

func NewTower(name string, numDisks int) *Tower {
	tower := &Tower{
		Name:  name,
		Disks: arraystack.New(),
	}

	tower.Disks.Clear()

	for i := numDisks; i > 0; i-- {
		tower.Disks.Push(i)
	}

	return tower
}

func (t *Tower) TakeTopFrom(tower *Tower) {
	if tower.Disks.Empty() {
		return
	}

	value, _ := tower.Disks.Pop()
	t.Disks.Push(value)
}

// TowerOfHanoi - решение задачи Ханойской башни
//
// Диск 1: A → C
// Диск 2: A → B
// Диск 1: C → B
// Диск 3: A → C
// Диск 1: B → A
// Диск 2: B → C
// Диск 1: A → C
//
// В классическом алгоритме Ханойской башни порядок аргументов всегда должен соответствовать логике "откуда, куда, через что"
func TowerOfHanoi(n int, origin, destination, auxiliary *Tower) {
	if n == 1 {
		destination.TakeTopFrom(origin)

		return
	}

	// Шаг 1: переместить n-1 дисков с origin на auxiliary (через destination)
	TowerOfHanoi(n-1, origin, auxiliary, destination)

	// Шаг 2: переместить самый большой диск с origin на destination
	destination.TakeTopFrom(origin)

	// Шаг 3: переместить n-1 дисков с auxiliary на destination (через origin)
	TowerOfHanoi(n-1, auxiliary, destination, origin)
}
