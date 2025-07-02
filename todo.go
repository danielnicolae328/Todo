package main

import (
	"fmt"
	"time"
)

type todo struct {
	title     string
	createdAt time.Time
}

var listOfTodos []todo

func createTodo() {

	var todoTitle string

	fmt.Print("Enter Todo title:")
	fmt.Scan(&todoTitle)

	todoTask := todo{title: todoTitle, createdAt: time.Now()}

	listOfTodos = append(listOfTodos, todoTask)
	fmt.Println(listOfTodos)

}
