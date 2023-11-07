package main

import (
	"fmt"
	"time"
)

/*
func main() {
	go goroutine()              // Начало горутины
	time.Sleep(4 * time.Second) // Ожидание выполнения функций
} // Здесь все горутины останавливаются
func goroutine() {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello")
}*/

/*
func main() {
	// небуферизованный канал int-ов
	ic := make(chan int)
	// буферизованный канал на 10 строк
	sc := make(chan string, 10)


	value := <-ic
	ic <- value
	close(ic)

	fmt.Println(value)
	for i := 0; i < 5; i++ {
		go goroutine(i)
	}
	time.Sleep(4 * time.Second)
}

func goroutine(id int) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello ", id)
}
*/

/*
func producer(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
}
func consumer(ch <-chan int) {
	for num := range ch {
		fmt.Println("Получено число:", num)
	}
}
func main() {
	ch := make(chan int)
	go producer(ch)
	consumer(ch)
}*/

func main() {
	c := make(chan int) // Делает канал для связи
	for i := 0; i < 5; i++ {
		go goroutine(i, c)
	}
	for i := 0; i < 5; i++ {
		goroutineID := <-c // Получает значение от канала
		fmt.Println("goroutine ", goroutineID, " has finished sleeping")
	}
}

func goroutine(id int, c chan int) { // Объявляет канал как аргумент
	time.Sleep(3 * time.Second)
	fmt.Println("Hello ", id)
	c <- id // Отправляет значение обратно к main
}
