type Todo struct {
	Title     string
	Completed bool
	//Root      js.Value
	tree *vdom.Tree
}

type TodoList struct {
	Todos []Todo
	Component
}

type Component struct {
	Name     string
	Root     js.Value
	Tree     *vdom.Tree
	Template string
}
