package locks

type Locker struct {
	chLock chan bool // Булево занимает меньше памяти чем пустая структура, хоть это и не очевидно
}

func NewLocker() *Locker {
	return &Locker{
		chLock: make(chan bool), // size 0 by default
	}
}

func (l *Locker) lock() {
	l.chLock <- true
}

func (l *Locker) unlock() {
	<-l.chLock
}

func (l *Locker) rLock() {
	// Реализация упрощена
}

func (l *Locker) rUnlock() {
	// Реализация упрощена
}

// Реализация rlock для правильной ждущей семантики, т.е.
// Lock -> ждет пока все читатели не закончат чтение
// RLock -> ждет пока не разблокируется lock и добавляет себя как +1 читателя
// RUnlock -> убирает одного читателя
// Unlock -> снимает блокировку с других горутин, читателей и писателей
//
// Такая семантика требует использования атомик типа данных Int
// Поэтому я решил не реализовывать это в рамках задачи про каналы

// CHANNEL
type SomeStruct struct {
	nums   []int
	locker *Locker
}

func NewSomeStruct(size int) *SomeStruct {
	return &SomeStruct{
		nums:   make([]int, 0, size),
		locker: NewLocker(),
	}
}

func (s *SomeStruct) UnsafeAppend(v int) {
	s.nums = append(s.nums, v)
}

func (s *SomeStruct) SafeAppend(v int) {
	s.locker.lock()
	defer s.locker.unlock()

	s.nums = append(s.nums, v)
}

func (s *SomeStruct) UnsafeRead(idx int) int {
	return s.nums[idx]
}

func (s *SomeStruct) SafeRead(idx int) int {
	s.locker.rLock()
	defer s.locker.rUnlock()

	return s.nums[idx]
}
