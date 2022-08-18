package main

import (
	"fmt"
	"time"
)

type MyError struct {
	message     string
	createdTime time.Time
}

func (m MyError) Error() string {
	return fmt.Sprintf("%s. Когда: %s", m.message, m.createdTime)
}

func fork(a [3]int) (err error) {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("Паника: ", v)
			err = MyError{message: "Ошибочка вышла", createdTime: time.Now()}
		}
	}()
	for i := 0; i < len(a)+1; i++ {
		fmt.Println(a[i])
	}
	return nil
}

func main() {
	s := [3]int{1, 2, 3}

	err := fork(s)
	if err != nil {
		fmt.Println(err)
	}
}
