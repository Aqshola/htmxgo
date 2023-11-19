package types

type Todo struct {
	Id    string
	Value string
	Done  bool
}

type ListTodo []Todo
