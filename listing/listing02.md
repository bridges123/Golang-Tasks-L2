Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
Вывод: 2 1
defer func выполняются в порядке LIFO, после выполнения всей функции.
То есть в test() сначала x = 1, затем в defer функции x увеличивается до 2, и возвращается эта увеличенная переменная
В случае anotherTest() сначала объявляется x, затем x = 1, после чего внутри defer func х увеличивается, но возвращается изначально объявленный x = 1
```
