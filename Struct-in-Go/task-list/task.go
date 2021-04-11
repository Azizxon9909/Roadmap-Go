package main

import (
	"fmt"
	"time"
)

type Task struct {
	title string
	task  string
	employee string
	deadline string
}

func (t Task) getTask() string {
	return "Title:"+t.title + " Task:" + t.task+ " Employee:" + t.employee+ " Deadline:" + t.deadline
}
func (t *Task) updateTask(title, task, employee, deadline string) {
	t.title = title
	t.task = task
	t.employee = employee
	t.deadline = deadline
}
func main() {

	c1 := Task{"Golang", "To-Do app", "Azizxon", time.April.String()}
	c2 := Task{"NodeJs", "RESTfull api", "Jahongir", time.December.String()}

	fmt.Println(c1.getTask())
	fmt.Println(c2.getTask())

	c2.updateTask("Python", "e-Commerce", "Fazliddin", time.July.String())
	fmt.Println(c2.getTask())

}