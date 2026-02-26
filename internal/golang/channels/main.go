package main

import (
	"fmt"
)

const bufn = 10

func main() {
	ch := make(chan int, bufn)

	printState(ch, "START")

	// Если сделать 11, получим дедлок. Блокировка мейн горутины
	for i := range bufn {
		write(ch, i+1)
	}

	printState(ch, "BEFORE CLOSE")

	close(ch)

	printState(ch, "CLOSE")

	// Пытаемся прочитать больше, чем позволяет буфер.
	// Нет дедлока. Потому, что после закрытия канала чтение из канала происходит без блокировки.
	// Если закомментировать close(ch) сверху, то будет дедлок.
	// Канал открыт и читатель ждет, пока кто-то запишет значение. Но этому не бывать.
	for range bufn + 100 {
		read(ch)
	}
}

func write(ch chan int, v int) {
	ch <- v

	printState(ch, fmt.Sprintf("sent v=%#v", v))
}

func read(ch chan int) {
	v, ok := <-ch
	printState(ch, fmt.Sprintf("read v=%#v ok=%#v", v, ok))
}

func printState(ch chan int, msg string) {
	fmt.Printf("%s\nch=%#v len=%d cap=%d\n\n", msg, ch, len(ch), cap(ch))
}
