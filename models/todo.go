package models

import "go_todo/dao"

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func AddTodo(todo *Todo) (err error) {
	return dao.DB.Create(&todo).Error
}

func DeleteTodo(todo *Todo) (err error) {
	return dao.DB.Delete(todo).Error
}

func UpdateTodo(todo *Todo) (err error) {
	return dao.DB.Save(todo).Error
}

func FindTodo(id string) (todo *Todo, err error) {
	todo = new(Todo)
	err = dao.DB.Where("id=?", id).First(todo).Error
	if err != nil {
		return nil, err
	}
	return
}

func FindTodoList() (todoList []*Todo, err error) {
	err = dao.DB.Find(&todoList).Error
	if err != nil {
		return nil, err
	}
	return
}
