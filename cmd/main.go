package main

import (
	"fmt"
	"strconv"
	"todo/src/model"
)

func main() {
	var i int
	for i != 4 {
		fmt.Println("Welcome to our Todo list!")
		fmt.Println("Choose from the menu:")
		fmt.Println("1 - View Todos\n2 - Add Todo\n3 - Remove Todo\n4 - Exit")
		fmt.Scan(&i)
		switch i {
		case 1:
			fmt.Println("Those are the todo items: ")
			fmt.Println(model.GetTodos())
		case 2:
			var name, desc string
			fmt.Println("Add Todo name: ")
			fmt.Scan(&name)
			fmt.Println("Add Todo description: ")
			fmt.Scan(&desc)
			newTodo := model.TodoFactory(name, desc)
			model.AddTodo(newTodo)
			fmt.Println("Todo added!")
		case 3:
			var id string
			fmt.Println("Removo todo by id: ")
			fmt.Scan(&id)
			intId, _ := strconv.ParseInt(id, 0, 32)
			model.RemoveTodo(int(intId))
		}
	}
	fmt.Println("\nYou're exiting the program. See you later!")
}
