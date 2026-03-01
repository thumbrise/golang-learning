package locks

import "sync"

// MUTEX реализация
type SomeStruct2 struct {
	nums []int
	mu   sync.RWMutex
}

func NewSomeStruct2(size int) *SomeStruct2 {
	return &SomeStruct2{
		nums: make([]int, 0, size),
	}
}

func (s *SomeStruct2) UnsafeAppend(v int) {
	s.nums = append(s.nums, v)
}

func (s *SomeStruct2) SafeAppend(v int) {
	// Заблокирует конкуретные записи
	// Также заблокирует читателей, если операция вызовет RLock, что проектируется на уровне публичных функций структуры
	s.mu.Lock()
	defer s.mu.Unlock()

	s.nums = append(s.nums, v)
}

func (s *SomeStruct2) UnsafeRead(idx int) int {
	return s.nums[idx]
}

func (s *SomeStruct2) SafeRead(idx int) int {
	// В случае блокировки читатель подождет, пока не закончится опасная операции записи.
	// Это позволяет:
	// 1 - не терять пропускную способность для чтения
	// 2 - не блокировать запись, ведь нам нужно только читать
	// 3 - не читать в моменты, когда читать опасно
	// Нельзя ни в коем случае делать вложенные блокировки Они никогда не разблокируются =) и будут ждать друг друга. Дедлок
	// Хорошей практикой является проектирование блокировок на уровне публичных методов как единного входа в логику
	// Если где-то нужна дополнительная блокировка, то следует использовать дополнительный мютекс или даже отдельную структуру
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.nums[idx]
}
