package handlers

import (
	"html/template"
	"htmxgo/data"
	"htmxgo/helpers"
	"htmxgo/types"
	"net/http"

	"github.com/google/uuid"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmplPath := helpers.GetTemplatePath("/templates/pages/todo.html")
	cardPath := helpers.GetTemplatePath("/templates/components/todolist.html")

	props := map[string]types.ListTodo{
		"Todos": data.DataTodoList,
	}

	tmpl := template.Must(template.ParseFiles(tmplPath, cardPath))
	tmpl.Execute(w, props)
}

func AddTodo(w http.ResponseWriter, r *http.Request) {
	/** FORM **/
	done := false
	value := r.PostFormValue("value")
	id, _ := uuid.NewRandom()

	cardPath := helpers.GetTemplatePath("/templates/components/todolist.html")

	data.DataTodoList = append(data.DataTodoList, types.Todo{
		Id:    id.String(),
		Value: value,
		Done:  done,
	})

	tmpl := template.Must(template.ParseFiles(cardPath))
	tmpl.ExecuteTemplate(w, "Todolist", types.Todo{
		Id:    id.String(),
		Value: value,
		Done:  done,
	})

}

func RemoveTodo(w http.ResponseWriter, r *http.Request) {
	id := helpers.GetField(r, 0)
	filteredTodo := types.ListTodo{}
	for _, el := range data.DataTodoList {
		if el.Id != id {
			filteredTodo = append(filteredTodo, el)
		}
	}
	data.DataTodoList = filteredTodo
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	id := helpers.GetField(r, 0)
	value := r.PostFormValue("done")
	parsedDone := false

	cardPath := helpers.GetTemplatePath("/templates/components/todolist.html")

	if value == "on" {
		parsedDone = true
	}

	idxUpdate := -1
	for idx, el := range data.DataTodoList {
		if el.Id == id {
			idxUpdate = idx
			break
		}
	}

	if idxUpdate != -1 {
		data.DataTodoList[idxUpdate].Done = parsedDone

		sendData := data.DataTodoList[idxUpdate]
		tmpl := template.Must(template.ParseFiles(cardPath))
		tmpl.ExecuteTemplate(w, "Todolist", types.Todo{
			Id:    sendData.Id,
			Value: sendData.Value,
			Done:  sendData.Done,
		})
	}

}
